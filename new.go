package logr

import (
	"github.com/go-jarvis/logr/slogx"
	"golang.org/x/exp/slog"
)

func Default() Logger {
	return &levelLogger{
		slog:  slogx.DefaultLogger(),
		level: InfoLevel,
	}
}

type Config struct {
	Level   string
	SLogger *slog.Logger
}

func New(c Config) Logger {
	return &levelLogger{
		level: LevelFromText(c.Level),
		slog:  c.SLogger,
	}
}
