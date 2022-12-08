package slogx

import "golang.org/x/exp/slog"

func DefaultLogger() *slog.Logger {
	return slog.New(DefaultHandler())
}
