package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	jsoniter "github.com/json-iterator/go"
	heimdallapp "github.com/maticnetwork/heimdall/app"
	authTypes "github.com/maticnetwork/heimdall/auth/types"
	borTypes "github.com/maticnetwork/heimdall/bor/types"
	"github.com/maticnetwork/heimdall/cmd/heimdalld/service"
	"github.com/maticnetwork/heimdall/helper"
	slashingTypes "github.com/maticnetwork/heimdall/slashing/types"
	stakingTypes "github.com/maticnetwork/heimdall/staking/types"
	topupTypes "github.com/maticnetwork/heimdall/topup/types"
	hmTypes "github.com/maticnetwork/heimdall/types"
	"github.com/tendermint/iavl/common"
	tmcfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	tmTypes "github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
	db "github.com/tendermint/tm-db"
	"golang.org/x/sync/errgroup"
)

var logger = helper.Logger.With("module", "heimdall/app")

var DefaultServiceCfg = HeimdallServiceCfg{
	PruningStrategy: "syncable",
}

type HeimdallServiceCfg struct {
	Chain           string
	DataDir         string
	Seeds           string
	PruningStrategy string //"Pruning strategy: syncable, nothing, everything"
}

func Start(ctx context.Context, cfg *HeimdallServiceCfg) error {
	serverCtx := server.NewDefaultContext()

	home := filepath.Join(cfg.DataDir, "heimdall")

	serverCtx.Config.RootDir = home
	serverCtx.Config.P2P.RootDir = filepath.Join(home, "p2p")

	// initialize heimdall if needed (do not force!)
	initConfig := &initConfig{
		chainID:     "", // chain id should be auto generated if chain flag is not set to mumbai or mainnet
		chain:       cfg.Chain,
		validatorID: 1, // default id for validator
		clientHome:  home,
		forceInit:   false,
	}

	cdc := heimdallapp.MakeCodec()

	if err := initConfig.heimdallInit(cdc, serverCtx.Config); err != nil {
		return fmt.Errorf("failed init heimdall: %s", err)
	}

	db, err := openDB(home)
	if err != nil {
		return fmt.Errorf("failed to open DB: %s", err)
	}

	// init heimdall config
	helper.InitHeimdallConfig(home)

	seedsFlagValue := cfg.Seeds

	if seedsFlagValue != "" {
		serverCtx.Config.P2P.Seeds = seedsFlagValue
	}

	if serverCtx.Config.P2P.Seeds == "" {
		switch helper.GetConfig().Chain {
		case helper.MainChain:
			serverCtx.Config.P2P.Seeds = helper.DefaultMainnetSeeds
		case helper.MumbaiChain:
			serverCtx.Config.P2P.Seeds = helper.DefaultTestnetSeeds
		}
	}

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	//app := heimdallapp.NewHeimdallApp(logger, db,
	//	baseapp.SetPruning(store.NewPruningOptionsFromString(cfg.PruningStrategy)))

	app := NewHeimdallConsumer(logger, db,
		baseapp.SetPruning(store.NewPruningOptionsFromString(cfg.PruningStrategy)))

	nodeKey, err := p2p.LoadOrGenNodeKey(serverCtx.Config.NodeKeyFile())

	if err != nil {
		return fmt.Errorf("failed to load or gen node key: %s", err)
	}

	server.UpgradeOldPrivValFile(serverCtx.Config)

	// create & start tendermint node
	tmNode, err := node.NewNode(
		serverCtx.Config,
		privval.LoadOrGenFilePV(serverCtx.Config.PrivValidatorKeyFile(), serverCtx.Config.PrivValidatorStateFile()),
		nodeKey,
		proxy.NewLocalClientCreator(app),
		node.DefaultGenesisDocProviderFunc(serverCtx.Config),
		node.DefaultDBProvider,
		node.DefaultMetricsProvider(serverCtx.Config.Instrumentation),
		serverCtx.Logger.With("module", "node"),
	)
	if err != nil {
		return fmt.Errorf("failed to create new node: %s", err)
	}

	// start Tendermint node here
	if err = tmNode.Start(); err != nil {
		return fmt.Errorf("failed to start Tendermint node: %s", err)
	}

	// using group context makes sense in case that if one of
	// the processes produces error the rest will go and shutdown
	g, gCtx := errgroup.WithContext(ctx)

	// stop phase for Tendermint node
	g.Go(func() error {
		// wait here for interrupt signal or
		// until something in the group returns non-nil error
		<-gCtx.Done()
		serverCtx.Logger.Info("exiting...")

		if tmNode.IsRunning() {
			return tmNode.Stop()
		}

		db.Close()

		return nil
	})

	// wait here for all go routines to finish,
	// or something to break
	if err := g.Wait(); err != nil {
		serverCtx.Logger.Error("Error shutting down services", "Error", err)
		return err
	}

	return nil
}

func openDB(rootDir string) (db.DB, error) {
	/*  TODO init MDBX based DB - home should be data dir
	dataDir := filepath.Join(rootDir, "data")
	return sdk.NewLevelDB("application", dataDir)
	*/
	return db.NewMemDB(), nil
}

type initConfig struct {
	clientHome  string
	chainID     string
	validatorID int64
	chain       string
	forceInit   bool
}

func (cfg initConfig) heimdallInit(cdc *codec.Codec, config *tmcfg.Config) error {
	configDir := filepath.Join(config.RootDir, "config")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(config.P2P.RootDir, "config"), 0755); err != nil {
		return err
	}

	conf := helper.GetDefaultHeimdallConfig()
	conf.Chain = cfg.chain

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	service.WriteDefaultHeimdallConfig(filepath.Join(configDir, "heimdall-config.toml"), conf)

	nodeID, valPubKey, _, err := service.InitializeNodeValidatorFiles(config)
	if err != nil {
		return err
	}

	// do not execute init if forceInit is false and genesis.json already exists (or we do not have permission to write to file)
	writeGenesis := cfg.forceInit

	if !writeGenesis {
		// When not forcing, check if genesis file exists
		_, err := os.Stat(config.GenesisFile())
		if err != nil && errors.Is(err, os.ErrNotExist) {
			logger.Info(fmt.Sprintf("Genesis file %v not found, writing genesis file\n", config.GenesisFile()))

			writeGenesis = true
		} else if err == nil {
			logger.Info(fmt.Sprintf("Found genesis file %v, skipping writing genesis file\n", config.GenesisFile()))
		} else {
			logger.Error(fmt.Sprintf("Error checking if genesis file %v exists: %v\n", config.GenesisFile(), err))
			return err
		}
	} else {
		logger.Info(fmt.Sprintf("Force writing genesis file to %v\n", config.GenesisFile()))
	}

	if writeGenesis {
		genesisCreated, err := helper.WriteGenesisFile(cfg.chain, config.GenesisFile(), cdc)
		if err != nil {
			return err
		} else if genesisCreated {
			return nil
		}
	} else {
		return nil
	}

	// create chain id
	chainID := cfg.chainID
	if chainID == "" {
		chainID = fmt.Sprintf("heimdall-%v", common.RandStr(6))
	}

	// get pubkey
	newPubkey := service.CryptoKeyToPubkey(valPubKey)

	// create validator account
	validator := hmTypes.NewValidator(hmTypes.NewValidatorID(uint64(cfg.validatorID)),
		0, 0, 1, 1, newPubkey,
		hmTypes.BytesToHeimdallAddress(valPubKey.Address().Bytes()))

	// create dividend account for validator
	dividendAccount := hmTypes.NewDividendAccount(validator.Signer, service.ZeroIntString)

	vals := []*hmTypes.Validator{validator}
	validatorSet := hmTypes.NewValidatorSet(vals)

	dividendAccounts := []hmTypes.DividendAccount{dividendAccount}

	// create validator signing info
	valSigningInfo := hmTypes.NewValidatorSigningInfo(validator.ID, 0, 0, 0)
	valSigningInfoMap := make(map[string]hmTypes.ValidatorSigningInfo)
	valSigningInfoMap[valSigningInfo.ValID.String()] = valSigningInfo

	// create genesis state
	appStateBytes := heimdallapp.NewDefaultGenesisState()

	// auth state change
	appStateBytes, err = authTypes.SetGenesisStateToAppState(
		appStateBytes,
		[]authTypes.GenesisAccount{getGenesisAccount(validator.Signer.Bytes())},
	)
	if err != nil {
		return err
	}

	// staking state change
	appStateBytes, err = stakingTypes.SetGenesisStateToAppState(appStateBytes, vals, *validatorSet)
	if err != nil {
		return err
	}

	// slashing state change
	appStateBytes, err = slashingTypes.SetGenesisStateToAppState(appStateBytes, valSigningInfoMap)
	if err != nil {
		return err
	}

	// bor state change
	appStateBytes, err = borTypes.SetGenesisStateToAppState(appStateBytes, *validatorSet)
	if err != nil {
		return err
	}

	// topup state change
	appStateBytes, err = topupTypes.SetGenesisStateToAppState(appStateBytes, dividendAccounts)
	if err != nil {
		return err
	}

	// app state json
	appStateJSON, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(appStateBytes)
	if err != nil {
		return err
	}

	toPrint := struct {
		ChainID string `json:"chain_id"`
		NodeID  string `json:"node_id"`
	}{
		chainID,
		nodeID,
	}

	out, err := codec.MarshalJSONIndent(cdc, toPrint)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "%s\n", string(out))

	return writeGenesisFile(tmtime.Now(), config.GenesisFile(), chainID, appStateJSON)
}

func getGenesisAccount(address []byte) authTypes.GenesisAccount {
	acc := authTypes.NewBaseAccountWithAddress(hmTypes.BytesToHeimdallAddress(address))

	genesisBalance, _ := big.NewInt(0).SetString("1000000000000000000000", 10)

	if err := acc.SetCoins(sdk.Coins{sdk.Coin{Denom: authTypes.FeeToken, Amount: sdk.NewIntFromBigInt(genesisBalance)}}); err != nil {
		logger.Error("getGenesisAccount | SetCoins", "Error", err)
	}

	result, _ := authTypes.NewGenesisAccountI(&acc)

	return result
}

func writeGenesisFile(genesisTime time.Time, genesisFile, chainID string, appState json.RawMessage) error {
	genDoc := tmTypes.GenesisDoc{
		GenesisTime: genesisTime,
		ChainID:     chainID,
		AppState:    appState,
	}

	if genDoc.GenesisTime.IsZero() {
		genDoc.GenesisTime = tmtime.Now()
	}

	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}

	return genDoc.SaveAs(genesisFile)
}
