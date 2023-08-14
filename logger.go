package logr

import (
	"context"
)

type Logger interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(err error)
	Error(err error)

	With(args ...any) Logger

	// 启动计时器
	// log = log.Start()
	// defer log.Stop()
	Start() Logger
	Stop()

	Enabled(level Level) bool
	SetLevel(level Level) Logger

	// WithContext(context.Context) Logger
	// Context() context.Context
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
