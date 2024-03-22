package types

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"math/bits"

	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/rlp"
)

type BlobSidecars []*BlobSidecar

// Len returns the length of s.
func (s BlobSidecars) Len() int { return len(s) }

// EncodeIndex encodes the i'th BlobTxSidecar to w. Note that this does not check for errors
// because we assume that BlobSidecars will only ever contain valid sidecars
func (s BlobSidecars) EncodeIndex(i int, w *bytes.Buffer) {
	rlp.Encode(w, s[i])
}

type BlobSidecar struct {
	BlobTxSidecar
	BlockNumber *big.Int       `json:"blockNumber"`
	BlockHash   libcommon.Hash `json:"blockHash"`
	TxIndex     uint64         `json:"transactionIndex"`
	TxHash      libcommon.Hash `json:"transactionHash"`
}

func NewBlobSidecarFromTx(tx *BlobTxWrapper) *BlobSidecar {
	if len(tx.Blobs) == 0 {
		return nil
	}
	return &BlobSidecar{
		BlobTxSidecar: *tx.BlobTxSidecar(),
		TxHash:        tx.Hash(),
	}
}

func (s *BlobSidecar) SanityCheck(blockNumber *big.Int, blockHash libcommon.Hash) error {
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

// generate encode and decode rlp for BlobSidecar
func (s *BlobSidecar) EncodeRLP(w io.Writer) error {
	if err := rlp.Encode(w, s.BlobTxSidecar); err != nil {
		return err
	}
	if err := rlp.Encode(w, s.BlockNumber); err != nil {
		return err
	}
	if err := rlp.Encode(w, s.BlockHash); err != nil {
		return err
	}
	if err := rlp.Encode(w, s.TxIndex); err != nil {
		return err
	}
	if err := rlp.Encode(w, s.TxHash); err != nil {
		return err
	}
	return nil
}

// DecodeRLP decodes a BlobSidecar from an RLP stream.
func (sc *BlobSidecar) DecodeRLP(s *rlp.Stream) error {
	_, err := s.List()
	if err != nil {
		return err
	}
	if err := sc.BlobTxSidecar.DecodeRLP(s); err != nil {
		return err
	}
	var b []byte
	if b, err = s.Bytes(); err != nil {
		return err
	}
	sc.BlockNumber = new(big.Int).SetBytes(b)
	if b, err = s.Bytes(); err != nil {
		return err
	}
	sc.BlockHash = libcommon.BytesToHash(b)

	if sc.TxIndex, err = s.Uint(); err != nil {
		return err
	}

	if b, err = s.Bytes(); err != nil {
		return err
	}
	sc.TxHash = libcommon.BytesToHash(b)
	return s.ListEnd()
}

func (s *BlobSidecar) payloadSize() int {
	size := 0
	size += s.BlobTxSidecar.payloadSize()
	size += 1 + rlp.BigIntLenExcludingHead(s.BlockNumber)
	size += 32
	size += 8
	size += 32
	return size
}

func (s *BlobSidecar) EncodingSize() int {
	return 1 + s.payloadSize()
}

// BlobTxSidecars encode and decode rlp methods
func (scs BlobSidecars) EncodeRLP(w io.Writer) error {
	// prefix
	payloadSize := scs.payloadSize()
	b := make([]byte, 9)
	if err := EncodeStructSizePrefix(payloadSize, w, b); err != nil {
		return err
	}

	// sidecars
	for _, sc := range scs {
		if err := sc.EncodeRLP(w); err != nil {
			return err
		}
	}

	return nil
}

func (scs *BlobSidecars) DecodeRLP(s *rlp.Stream) error {
	_, err := s.List()
	if err != nil {
		return fmt.Errorf("open BlobSidecar: %w", err)
	}

	var sc BlobSidecar
	for {
		err := sc.DecodeRLP(s)
		if err != nil {
			return fmt.Errorf("decode BlobSidecar: %w", err)
		}
		*scs = append(*scs, &sc)
		if err = s.ListEnd(); err != nil {
			return nil
		}
	}
}

func (scs BlobSidecars) payloadSize() int {
	size := 1                                                   // 0xc7..0xcf
	size += libcommon.BitLenToByteLen(bits.Len(uint(len(scs)))) // length encoding size
	for _, sc := range scs {
		size += sc.payloadSize()
	}
	return size
}

func (scs BlobSidecars) EncodingSize() int {
	size := 9
	for _, sc := range scs {
		size += sc.EncodingSize()
	}
	return size
}
