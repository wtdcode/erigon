package bridge

import (
	"context"
	"fmt"

	"github.com/ledgerwatch/log/v3"

	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon/node"
	"github.com/ledgerwatch/erigon/node/nodecfg"
)

type Bridge struct {
	DB kv.RwDB
}

func NewBridge(ctx context.Context, config *nodecfg.Config, name string, readonly bool, logger log.Logger) (*Bridge, error) {
	// create new db
	db, err := node.OpenDatabase(ctx, config, kv.PolygonDB, name, readonly, logger)
	if err != nil {
		return nil, err
	}

	return &Bridge{
		DB: db,
	}, nil
}

func (b *Bridge) Run() {
	// start syncing
	fmt.Println("bridge is running")
}

func (b *Bridge) Close() {
	b.DB.Close()
}
