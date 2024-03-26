package core

import (
	"fmt"

	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/consensus"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/params"
)

// IsDataAvailable it checks that the blobTx block has available blob data
func IsDataAvailable(chain consensus.ChainHeaderReader, header *types.Header, body *types.RawBody) (err error) {
	if !chain.Config().IsCancun(header.Number.Uint64(), header.Time) {
		return nil
	}

	current := chain.CurrentHeader()
	if header.Number.Uint64()+params.MinBlocksForBlobRequests < current.Number.Uint64() {
		// if we needn't check DA of this block, just clean it
		body.CleanSidecars()
		return nil
	}

	// alloc block's versionedHashes
	sidecars := body.Sidecars
	versionedHashes := make([][]common.Hash, 0, len(body.Transactions))
	if len(versionedHashes) != len(sidecars) {
		return fmt.Errorf("blob info mismatch: sidecars %d, versionedHashes:%d", len(sidecars), len(versionedHashes))
	}
	blobIndex := 0
	txs, err := types.DecodeTransactions(body.Transactions)
	if err != nil {
		return err
	}
	for _, tx := range txs {
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
