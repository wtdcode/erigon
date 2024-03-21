package disk

import (
	"context"
	"time"

	"github.com/ledgerwatch/log/v3"
)

func UpdateDiskStats(ctx context.Context, logger log.Logger) {
	logEvery := time.NewTicker(5 * time.Second)
	defer logEvery.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Debug("[dbg] exit1")
			return
		case <-logEvery.C:

			log.Debug("[dbg] no exit1")
			if err := UpdatePrometheusDiskStats(); err != nil {
				logger.Warn("[disk] error disk fault stats", "err", err)
			}
		}
	}
}
