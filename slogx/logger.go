package slogx

import (
	"os"

	"golang.org/x/exp/slog"
)

func DefaultLogger() *slog.Logger {
	return slog.New(DefaultHandler())
}

func DefaultJsonLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout))
}
