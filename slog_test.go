package logr

import (
	"errors"
	"testing"

	"golang.org/x/exp/slog"
)

func TestDefault(t *testing.T) {

	log := Default().SetLevel(slog.DebugLevel)
	err := errors.New("New_ERROR")

	log.Debug("number=%d", 1)
	log.Info("number=%d", 1)
	log.Warn(err)
	log.Error(err)
}
