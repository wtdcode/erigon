package types

import (
	"bytes"
	"errors"
	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/rlp"
	"math/big"
)

type BlobSidecars []*BlobSidecar

func (s BlobSidecars) Len() int { return len(s) }

// EncodeIndex encodes the i'th BlobTxSidecar to w. Note that this does not check for errors
// because we assume that BlobSidecars will only ever contain valid sidecars
func (s BlobSidecars) EncodeIndex(i int, w *bytes.Buffer) {
	rlp.Encode(w, s[i])
}

type BlobSidecar struct {
	BlobTxSidecar
	BlockNumber *big.Int    `json:"blockNumber"`
	BlockHash   common.Hash `json:"blockHash"`
	TxIndex     uint64      `json:"transactionIndex"`
	TxHash      common.Hash `json:"transactionHash"`
}

func NewBlobSidecarFromTx(tx *Transaction) *BlobSidecar {
	return nil
}

func (s *BlobSidecar) SanityCheck(blockNumber *big.Int, blockHash common.Hash) error {
	if s.BlockNumber.Cmp(blockNumber) != 0 {
		return errors.New("BlobSidecar with wrong block number")
	}
	if s.BlockHash != blockHash {
		return errors.New("BlobSidecar with wrong block hash")
	}
	if len(s.Blobs) != len(s.Commitments) {
		return errors.New("BlobSidecar has wrong commitment length")
	}
	if len(s.Blobs) != len(s.Proofs) {
		return errors.New("BlobSidecar has wrong proof length")
	}
	return nil
}
