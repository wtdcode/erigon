package hexutil

import "github.com/ledgerwatch/erigon-lib/common/hexutility"

// This file is included to allow replace github.com/ethereum/go-ethereum => ./
// include here any methods which need to be redirected

func Encode(b []byte) string { return hexutility.Encode(b) }
