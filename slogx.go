package logr

import (
	"fmt"

	"github.com/go-jarvis/logr/slogx"
	"golang.org/x/exp/slog"
)

var _ Logger = &logger{}

type logger struct {
	slog  *slog.Logger
	level slog.Level

	args []any
}

func Default() *logger {
	return &logger{
		slog:  slogx.Default(),
		level: slog.InfoLevel,
	}
}

func (log *logger) Debug(msg string, args ...any) {
	if log.Enabled(slog.DebugLevel) {
		log.slog.LogDepth(0, slog.DebugLevel, fmt.Sprintf(msg, args...))
	}
}

func (log *logger) Info(msg string, args ...any) {
	if log.Enabled(slog.InfoLevel) {
		log.slog.LogDepth(0, slog.InfoLevel, fmt.Sprintf(msg, args...))
	}
}

func (log *logger) Warn(err error) {
	if log.Enabled(slog.WarnLevel) {
		log.slog.LogDepth(0, slog.WarnLevel, err.Error())
	}
}

func (log *logger) Error(err error) {
	if log.Enabled(slog.ErrorLevel) {
		log.slog.LogDepth(0, slog.ErrorLevel, err.Error())
	}
}

func (log *logger) With(args ...any) Logger {
	if len(args)%2 != 0 {
		args = append(args, "Unknown_LACK")
	}
	if log.args == nil {
		log.args = make([]any, 0)
	}

	log.args = append(log.args, args...)

	return log
}

// Enabled return log level result
func (log *logger) Enabled(level slog.Level) bool {
	return log.level <= level
}

// SetLevel set level
func (log *logger) SetLevel(level slog.Level) Logger {
	return &logger{
		slog:  log.slog,
		level: level,
	}
}
