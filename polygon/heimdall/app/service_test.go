package app

import (
	"context"
	"testing"
)

func TestHeimdallServiceStart(t *testing.T) {

	cfg := DefaultServiceCfg
	cfg.Chain = "mumbai"
	cfg.DataDir = t.TempDir()

	Start(context.Background(), &cfg)
}
