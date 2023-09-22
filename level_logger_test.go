package logr

import (
	"testing"

	"github.com/go-jarvis/logr/slogx"
	"golang.org/x/exp/slog"
)

func TestLogger(t *testing.T) {
	// log := DefaultLevelLogger()
	log := New(
		Config{
			log: slogx.DefaultJSONLogger(),
			// log:   slogx.DefaultLogger(),
			level: DebugLevel,
		},
	)

	slog.Info("babaa")

	log.Debug("this is %s", "debug")
	log.Info("this is %s", "info")

	log = log.With("key1", "val1")
	log.Warn("this is %s", "warn")

	log = log.With("key2", "val2")
	log.Error("this is %s", "error")

	log.Fatal("this is %s", "fatal")
}
