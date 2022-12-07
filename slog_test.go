package logr

import (
	"errors"
	"testing"
	"time"

	"golang.org/x/exp/slog"
)

func TestDefault(t *testing.T) {

	log := Default().SetLevel(slog.DebugLevel)
	err := errors.New("New_ERROR")

	log.Debug("number=%d", 1)
	log.Info("number=%d", 1)
	log.Warn(err)
	log.Error(err)

	log = log.Start()
	defer log.Stop()

	time.Sleep(2 * time.Second)

}
