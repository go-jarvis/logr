package logr

import (
	"errors"
	"testing"
	"time"

	"golang.org/x/exp/slog"
)

func timeStamp() Valuer {
	return func() any {
		return time.Now().Format(time.RFC3339)
	}
}

func timeStamp2() any {
	return time.Now().Format(time.RFC3339)
}

func TestDefault(t *testing.T) {

	log := Default().SetLevel(slog.DebugLevel)
	err := errors.New("New_ERROR")

	log = log.With("timestamp", timeStamp(), "kk", "vv", "ts", timeStamp2)

	log.Debug("number=%d", 1)
	log.Info("number=%d", 1)
	log.Warn(err)
	log.Error(err)

	log = log.Start()
	defer log.Stop()

	time.Sleep(2 * time.Second)

}
