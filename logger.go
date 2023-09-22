package logr

import (
	"context"
)

type Logger interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)

	// With("key1", "val1", "key2", "val2")
	With(kvs ...any) Logger
}

type LoggerKey int

const defaultLoggerKey LoggerKey = 0

func WithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, defaultLoggerKey, log)
}

func FromContext(ctx context.Context) Logger {
	val := ctx.Value(defaultLoggerKey)
	if log, ok := val.(Logger); ok {
		return log
	}

	return nil
}
