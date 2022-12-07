package logr

import (
	"golang.org/x/exp/slog"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(err error)
	Error(err error)

	Enabled(level slog.Level) bool
	SetLevel(level slog.Level) Logger
}
