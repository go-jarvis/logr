package slogx

import (
	"os"

	"golang.org/x/exp/slog"
)

func DefaultLogger() *slog.Logger {
	return slog.New(
		newHandler(
			slog.Default().Handler(),
		),
	)

}

func DefaultJSONLogger() *slog.Logger {
	return slog.New(
		newHandler(
			slog.NewJSONHandler(os.Stdout, nil),
		),
	)
}
