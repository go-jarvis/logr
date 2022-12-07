package logr

import (
	"context"
	"errors"
	"testing"
	"time"
)

func ExampleDefault() {
	log := Default()
	log = log.With("k", "v")
	log.Info("example", "kk", "vv")
	// Output:
}

func TestDefaultJSON(t *testing.T) {
	log := DefaultJSON()
	log = log.With("k1", "v1")

	_, log2 := log.Start(context.Background())
	defer log2.Stop()

	err := errors.New("Test_ERR")

	log2.Info("baa")
	log2.Warn(err)

	time.Sleep(2 * time.Second)
	log2.Error(err)

	// Output:
}
