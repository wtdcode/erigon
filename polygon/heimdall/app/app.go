package app

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	jsoniter "github.com/json-iterator/go"
	heimdallapp "github.com/maticnetwork/heimdall/app"
	"github.com/maticnetwork/heimdall/auth"
	authTypes "github.com/maticnetwork/heimdall/auth/types"
	"github.com/maticnetwork/heimdall/bank"
	bankTypes "github.com/maticnetwork/heimdall/bank/types"
	"github.com/maticnetwork/heimdall/chainmanager"
	chainmanagerTypes "github.com/maticnetwork/heimdall/chainmanager/types"
	"github.com/maticnetwork/heimdall/checkpoint"
	checkpointTypes "github.com/maticnetwork/heimdall/checkpoint/types"
	"github.com/maticnetwork/heimdall/common"
	govTypes "github.com/maticnetwork/heimdall/gov/types"
	"github.com/maticnetwork/heimdall/helper"
	"github.com/maticnetwork/heimdall/params"
	"github.com/maticnetwork/heimdall/params/subspace"
	paramsTypes "github.com/maticnetwork/heimdall/params/types"
	"github.com/maticnetwork/heimdall/staking"
	stakingTypes "github.com/maticnetwork/heimdall/staking/types"
	"github.com/maticnetwork/heimdall/supply"
	supplyTypes "github.com/maticnetwork/heimdall/supply/types"
	"github.com/maticnetwork/heimdall/topup"
	"github.com/maticnetwork/heimdall/version"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

const (
	AppName = "HeimdallConsumer"
)

var (
	ModuleBasics = module.NewBasicManager(
		params.AppModuleBasic{},
		checkpoint.AppModuleBasic{},
		//borModuleBasic{},
		//clerkModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authTypes.FeeCollectorName: nil,
		govTypes.ModuleName:        {},
	}
)

// GenesisState the genesis state of the blockchain is represented here as a map of raw json messages keyed by an identifier string
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() GenesisState {
	return ModuleBasics.DefaultGenesis()
}

// HeimdallConsumer main heimdall app
type HeimdallConsumer struct {
	heimdallapp.HeimdallApp
	caller helper.ContractCaller
}

func mustField(val interface{}, field string) reflect.StructField {
	if f, ok := reflect.TypeOf(val).FieldByName(field); !ok {
		panic(fmt.Errorf("can't get field %s for %s", reflect.TypeOf(val), field))
	} else {
		return f
	}
}

var fields = struct {
	HeimdallApp reflect.StructField
	mm          reflect.StructField
	subspaces   reflect.StructField
	keys        reflect.StructField
	tkeys       reflect.StructField
}{
	HeimdallApp: mustField(HeimdallConsumer{}, "HeimdallApp"),
	mm:          mustField(heimdallapp.HeimdallApp{}, "mm"),
	subspaces:   mustField(heimdallapp.HeimdallApp{}, "subspaces"),
	keys:        mustField(heimdallapp.HeimdallApp{}, "keys"),
	tkeys:       mustField(heimdallapp.HeimdallApp{}, "tkeys"),
}

//go:linkname typedmemmove reflect.typedmemmove
func typedmemmove(rtype unsafe.Pointer, dst, src unsafe.Pointer)

// NewHeimdallConsumer creates heimdall app
func NewHeimdallConsumer(logger log.Logger, db dbm.DB, baseAppOptions ...func(*bam.BaseApp)) *HeimdallConsumer {
	type iface struct {
		_, data unsafe.Pointer
	}

	unpackIf := func(obj interface{}) *iface {
		return (*iface)(unsafe.Pointer(&obj))
	}

	// create and register app-level codec for TXs and accounts
	cdc := heimdallapp.MakeCodec()

	// set prefix
	config := sdk.GetConfig()
	config.Seal()

	// base app
	bApp := bam.NewBaseApp(AppName, logger, db, authTypes.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(nil)
	bApp.SetAppVersion(version.Version)

	// keys
	keys := sdk.NewKVStoreKeys(
		bam.MainStoreKey,
		authTypes.StoreKey,
		bankTypes.StoreKey,
		supplyTypes.StoreKey,
		chainmanagerTypes.StoreKey,
		stakingTypes.StoreKey,
		checkpointTypes.StoreKey,
		//borTypes.StoreKey,
		//clerkTypes.StoreKey,
		paramsTypes.StoreKey,
	)

	tkeys := sdk.NewTransientStoreKeys(paramsTypes.TStoreKey)

	// create heimdall app
	var app = &HeimdallConsumer{
		HeimdallApp: heimdallapp.HeimdallApp{
			BaseApp: bApp,
		},
	}

	app.SetCodec(cdc)

	fptr := unsafe.Pointer(uintptr(unsafe.Pointer(app)) + fields.HeimdallApp.Offset + fields.keys.Offset)
	typedmemmove(unpackIf(fields.keys.Type).data, fptr, unsafe.Pointer(&keys))

	fptr = unsafe.Pointer(uintptr(unsafe.Pointer(app)) + fields.HeimdallApp.Offset + fields.tkeys.Offset)
	typedmemmove(unpackIf(fields.tkeys.Type).data, fptr, unsafe.Pointer(&tkeys))

	// init params keeper and subspaces
	app.ParamsKeeper = params.NewKeeper(cdc, keys[paramsTypes.StoreKey], tkeys[paramsTypes.TStoreKey], paramsTypes.DefaultCodespace)

	subspaces := map[string]subspace.Subspace{}

	fptr = unsafe.Pointer(uintptr(unsafe.Pointer(app)) + fields.HeimdallApp.Offset + fields.subspaces.Offset)
	typedmemmove(unpackIf(fields.subspaces.Type).data, fptr, unsafe.Pointer(&subspaces))

	subspaces[authTypes.ModuleName] = app.ParamsKeeper.Subspace(authTypes.DefaultParamspace)
	subspaces[chainmanagerTypes.ModuleName] = app.ParamsKeeper.Subspace(chainmanagerTypes.DefaultParamspace)
	subspaces[bankTypes.ModuleName] = app.ParamsKeeper.Subspace(bankTypes.DefaultParamspace)
	subspaces[supplyTypes.ModuleName] = app.ParamsKeeper.Subspace(supplyTypes.DefaultParamspace)
	subspaces[checkpointTypes.ModuleName] = app.ParamsKeeper.Subspace(checkpointTypes.DefaultParamspace)
	subspaces[stakingTypes.ModuleName] = app.ParamsKeeper.Subspace(stakingTypes.DefaultParamspace)
	//app.subspaces[borTypes.ModuleName] = app.ParamsKeeper.Subspace(borTypes.DefaultParamspace)
	//app.subspaces[clerkTypes.ModuleName] = app.ParamsKeeper.Subspace(clerkTypes.DefaultParamspace)

	moduleCommunicator := heimdallapp.ModuleCommunicator{App: &app.HeimdallApp}

	// create chain keeper
	app.ChainKeeper = chainmanager.NewKeeper(
		cdc,
		keys[chainmanagerTypes.StoreKey], // target store
		subspaces[chainmanagerTypes.ModuleName],
		common.DefaultCodespace,
		app.caller,
	)

	app.StakingKeeper = staking.NewKeeper(
		cdc,
		keys[stakingTypes.StoreKey], // target store
		subspaces[stakingTypes.ModuleName],
		common.DefaultCodespace,
		app.ChainKeeper,
		moduleCommunicator,
	)

	app.AccountKeeper = auth.NewAccountKeeper(
		cdc,
		keys[authTypes.StoreKey], // target store
		subspaces[authTypes.ModuleName],
		authTypes.ProtoBaseAccount, // prototype
	)

	app.BankKeeper = bank.NewKeeper(
		cdc,
		keys[bankTypes.StoreKey], // target store
		subspaces[bankTypes.ModuleName],
		bankTypes.DefaultCodespace,
		app.AccountKeeper,
		moduleCommunicator,
	)

	// bank keeper
	app.SupplyKeeper = supply.NewKeeper(
		cdc,
		keys[supplyTypes.StoreKey], // target store
		subspaces[supplyTypes.ModuleName],
		maccPerms,
		app.AccountKeeper,
		app.BankKeeper,
	)

	app.CheckpointKeeper = checkpoint.NewKeeper(
		cdc,
		keys[checkpointTypes.StoreKey], // target store
		subspaces[checkpointTypes.ModuleName],
		common.DefaultCodespace,
		app.StakingKeeper,
		app.ChainKeeper,
		moduleCommunicator,
	)

	/*
		app.BorKeeper = NewBorKeeper(
			app.cdc,
			keys[borTypes.StoreKey], // target store
			app.subspaces[borTypes.ModuleName],
			common.DefaultCodespace,
			app.ChainKeeper,
			app.StakingKeeper,
			app.caller,
		)

		app.ClerkKeeper = NewClerkKeeper(
			app.cdc,
			keys[clerkTypes.StoreKey], // target store
			app.subspaces[clerkTypes.ModuleName],
			common.DefaultCodespace,
			app.ChainKeeper,
		)
	*/
	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	mm := module.NewManager(
		checkpoint.NewAppModule(app.CheckpointKeeper, app.StakingKeeper, topup.Keeper{}, nil),
		//NewBorAppModule(app.BorKeeper, &app.caller),
		//NewClerkModule(app.ClerkKeeper, &app.caller),
	)

	fptr = unsafe.Pointer(uintptr(unsafe.Pointer(app)) + fields.HeimdallApp.Offset + fields.mm.Offset)
	typedmemmove(unpackIf(fields.mm.Type).data, fptr, unsafe.Pointer(&mm))

	mm.SetOrderInitGenesis(
		checkpointTypes.ModuleName,
		//borTypes.ModuleName,
		//clerkTypes.ModuleName,
	)

	mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// mount the multistore and load the latest state
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	// perform initialization logic
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// load latest version
	err := app.LoadLatestVersion(keys[bam.MainStoreKey])
	if err != nil {
		cmn.Exit(err.Error())
	}

	app.Seal()

	return app
}

// InitChainer initializes chain
func (app *HeimdallConsumer) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState

	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	// get validator updates
	if err := ModuleBasics.ValidateGenesis(genesisState); err != nil {
		panic(err)
	}

	// check fee collector module account
	if moduleAcc := app.SupplyKeeper.GetModuleAccount(ctx, authTypes.FeeCollectorName); moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", authTypes.FeeCollectorName))
	}

	// init genesis
	app.GetModuleManager().InitGenesis(ctx, genesisState)

	stakingState := stakingTypes.GetGenesisStateFromAppState(genesisState)
	checkpointState := checkpointTypes.GetGenesisStateFromAppState(genesisState)

	// check if validator is current validator
	// add to val updates else skip
	var valUpdates []abci.ValidatorUpdate

	for _, validator := range stakingState.Validators {
		if validator.IsCurrentValidator(checkpointState.AckCount) {
			// convert to Validator Update
			updateVal := abci.ValidatorUpdate{
				Power:  validator.VotingPower,
				PubKey: validator.PubKey.ABCIPubKey(),
			}
			// Add validator to validator updated to be processed below
			valUpdates = append(valUpdates, updateVal)
		}
	}

	// TODO make sure old validtors dont go in validator updates ie deactivated validators have to be removed
	// update validators
	return abci.ResponseInitChain{
		// validator updates
		Validators: valUpdates,
	}
}

// BeginBlocker application updates every begin block
func (app *HeimdallConsumer) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.GetModuleManager().BeginBlock(ctx, req)
}

// EndBlocker executes on each end block
func (app *HeimdallConsumer) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	var tmValUpdates []abci.ValidatorUpdate
	/*
		// --- Start update to new validators
		currentValidatorSet := app.StakingKeeper.GetValidatorSet(ctx)
		allValidators := app.StakingKeeper.GetAllValidators(ctx)
		ackCount := app.CheckpointKeeper.GetACKCount(ctx)

		// get validator updates
		setUpdates := helper.GetUpdatedValidators(
			&currentValidatorSet, // pointer to current validator set -- UpdateValidators will modify it
			allValidators,        // All validators
			ackCount,             // ack count
		)

		if len(setUpdates) > 0 {
			// create new validator set
			if err := currentValidatorSet.UpdateWithChangeSet(setUpdates); err != nil {
				// return with nothing
				logger.Error("Unable to update current validator set", "Error", err)
				return abci.ResponseEndBlock{}
			}

			//Hardfork to remove the rotation of validator list on stake update
			if ctx.BlockHeight() < helper.GetAalborgHardForkHeight() {
				// increment proposer priority
				currentValidatorSet.IncrementProposerPriority(1)
			}

			// validator set change
			logger.Debug("[ENDBLOCK] Updated current validator set", "proposer", currentValidatorSet.GetProposer())

			// save set in store
			if err := app.StakingKeeper.UpdateValidatorSetInStore(ctx, currentValidatorSet); err != nil {
				// return with nothing
				logger.Error("Unable to update current validator set in state", "Error", err)
				return abci.ResponseEndBlock{}
			}

			// convert updates from map to array
			for _, v := range setUpdates {
				tmValUpdates = append(tmValUpdates, abci.ValidatorUpdate{
					Power:  v.VotingPower,
					PubKey: v.PubKey.ABCIPubKey(),
				})
			}
		}
	*/
	// end block
	app.GetModuleManager().EndBlock(ctx, req)

	// send validator updates to peppermint
	return abci.ResponseEndBlock{
		ValidatorUpdates: tmValUpdates,
	}
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}

	return dupMaccPerms
}
