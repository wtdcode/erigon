package core

import (
	"fmt"
	"sync"

	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/consensus"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/params"
)

// IsDataAvailable it checks that the blobTx block has available blob data
func IsDataAvailable(chain consensus.ChainHeaderReader, block *types.Block) (err error) {
	if !chain.Config().IsCancun(block.Number().Uint64(), block.Time()) {
		return nil
	}

	current := chain.CurrentHeader()
	if block.NumberU64()+params.MinBlocksForBlobRequests < current.Number.Uint64() {
		// if we needn't check DA of this block, just clean it
		block.CleanSidecars()
		return nil
	}

	// alloc block's versionedHashes
	sidecars := block.Sidecars()
	versionedHashes := make([][]common.Hash, 0, len(block.Transactions()))
	if len(versionedHashes) != len(sidecars) {
		return fmt.Errorf("blob info mismatch: sidecars %d, versionedHashes:%d", len(sidecars), len(versionedHashes))
	}
	blobIndex := 0
	for _, tx := range block.Transactions() {
		hashes := tx.GetBlobHashes()
		if hashes == nil {
			continue
		}
		if err := sidecars[blobIndex].ValidateBlobTxSidecar(hashes); err != nil {
			return err
		}
		blobIndex++
	}

	return nil
}

func CheckDataAvailableInBatch(chainReader consensus.ChainHeaderReader, chain types.Blocks) (int, error) {
	if len(chain) == 1 {
		return 0, IsDataAvailable(chainReader, chain[0])
	}

	var (
		wg   sync.WaitGroup
		errs sync.Map
	)

	for i := range chain {
		wg.Add(1)
		func(index int, block *types.Block) {
			go func() {
				defer wg.Done()
				errs.Store(index, IsDataAvailable(chainReader, block))
			}()
		}(i, chain[i])
	}

	wg.Wait()
	for i := range chain {
		val, exist := errs.Load(i)
		if !exist || val == nil {
			continue
		}
		err := val.(error)
		if err != nil {
			return i, err
		}
	}
	return 0, nil
}
