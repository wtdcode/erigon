package gsubmit

import (
	"context"

	"github.com/ledgerwatch/erigon-lib/gointerfaces/sentinel"
)

type SubmitState struct {

	// the gossip handler
	sentinel sentinel.SentinelClient

	// if this is not empty, it means we are aggregating on that subnet
	isAggregating string
}

func (s *SubmitState) publishGossip(ctx context.Context) {
	s.sentinel.PublishGossip(ctx, &sentinel.GossipData{
		Data: nil,
		Type: sentinel.GossipType_AttesterSlashingGossipType,
	})

}
