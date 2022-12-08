package logr

import (
	"context"

	"golang.org/x/exp/slog"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(err error)
	Error(err error)

	With(args ...any) Logger

	Start() Logger
	Stop()

	Enabled(level slog.Level) bool
	SetLevel(level slog.Level) Logger

	WithContext(context.Context) Logger
	Context() context.Context
}

type LogrKey int

const defaultLogrKey LogrKey = 0

func WithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, defaultLogrKey, log)
}

func FromContext(ctx context.Context) Logger {
	val := ctx.Value(defaultLogrKey)
	if log, ok := val.(Logger); ok {
		return log
	}

	return Default()
}
