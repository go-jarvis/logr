package logr

import (
	"golang.org/x/exp/slog"
)

type Config struct {
	log   *slog.Logger
	level Level
}

func New(c Config) Logger {

	return &levelLogger{
		logger: c.log,
		level:  c.level,
	}
}
