package logr

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-jarvis/logr/slogx"
)

func newLogger() Logger {
	c := Config{
		Level:   "debug",
		SLogger: slogx.DefaultJsonLogger(),
	}
	return New(c)

}
func TestDefault(t *testing.T) {

	// log := Default().SetLevel(DebugLevel)
	log := newLogger()
	err := errors.New("New_ERROR")

	log = log.With(
		"kk", "vv",
		"caller", CallerFile(4, false),
		"gg",
	)

	log.Debug("number=%d", 1)
	log.Info("number=%d", 1)
	log.Warn(err)
	log.Error(err)

	ctx := WithLogger(context.Background(), log)
	subcaller(ctx)
}

func subcaller(ctx context.Context) {
	log := FromContext(ctx)

	log = log.Start() // time cost
	defer log.Stop()

	time.Sleep(532 * time.Millisecond)
	log.Info("account=%d", 100)
}
