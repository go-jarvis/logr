package slogx

import "golang.org/x/exp/slog"

func Default() *slog.Logger {
	return slog.New(defaultHandler())
}
