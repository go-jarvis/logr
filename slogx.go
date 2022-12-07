package logr

import (
	"fmt"
	"time"

	"github.com/go-jarvis/logr/slogx"
	"golang.org/x/exp/slog"
)

var _ Logger = &logger{}

type logger struct {
	slog  *slog.Logger
	level slog.Level

	kvs []any

	timer time.Time
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

func (log *logger) With(kvs ...any) Logger {
	if len(kvs)%2 != 0 {
		kvs = append(kvs, "Unknown_LACK")
	}
	if log.kvs == nil {
		log.kvs = make([]any, 0)
	}

	logc := log.copy()
	logc.slog = log.slog.With(kvs...)
	return logc
}

func (log *logger) Start() Logger {
	logc := log.copy()
	logc.timer = time.Now()

	return logc
}

func (log *logger) Stop() {
	cost := time.Now().Sub(log.timer).Milliseconds() / 1e3

	log.With("cost", cost).Info("time-cost")

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

func (log *logger) copy() *logger {
	return &logger{
		slog:  log.slog,
		level: log.level,
		kvs:   log.kvs,
	}
}
