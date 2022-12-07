package logr

import "context"

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(err error)
	Error(err error)

	// With params
	With(args ...any) Logger

	// Start timer
	// 		_, log2 := log.Start(context.Background())
	// 		defer log2.Stop()
	Start(ctx context.Context, args ...any) (context.Context, Logger)
	// Stop time, and print out cost in second
	Stop()
}

type LoggerType int

var loggerKey LoggerType = 0

func WithConetxt(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

func FromContext(ctx context.Context) Logger {
	val := ctx.Value(loggerKey)
	log, ok := val.(Logger)
	if !ok {
		return Default()
	}
	return log
}
