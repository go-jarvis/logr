package logr

import (
	"context"

	"github.com/go-jarvis/logr/slogx"
	"golang.org/x/exp/slog"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
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

type Config struct {
	level  string
	logger *slog.Logger
}

func New(c Config) Logger {
	if c.logger == nil {
		c.logger = slogx.DefaultLogger()
	}
	if c.level == "" {
		c.level = "info"
	}

	return &levelLogger{
		level: LevelFromText(c.level),
		slog:  c.logger,
	}
}
