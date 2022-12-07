package logr

import (
	"context"
	"testing"
	"time"
)

func TestValuer(t *testing.T) {

	log := DefaultJSON()

	log.With("timestamp", TimeStamp())
	log.With("caller", Caller(4, false))

	for i := 0; i < 1; i++ {
		log.Info("test-valuer")
		time.Sleep(1 * time.Second)
	}

	ctx := WithConetxt(context.Background(), log)
	_, log = log.Start(ctx, "start", "start-value")
	defer log.Stop()

	subcaller(ctx)
}

func subcaller(ctx context.Context) {
	log := FromContext(ctx)
	_, log = log.Start(ctx)
	defer log.Stop()

	time.Sleep(1 * time.Second)
	log.Info("subcaller")
}
