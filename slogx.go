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

	hasValuer bool
	kvs       []any

	timer time.Time
}

func Default() *logger {
	return &logger{
		slog:  slogx.Default(),
		level: slog.InfoLevel,
	}
}

// Log 绑定参数，打印日志
func (log *logger) Log(level slog.Level, msg string) {
	kvs := append([]any{}, log.kvs...)

	if log.hasValuer {
		kvs = bindValuer(kvs...)
	}

	log.slog.With(kvs...).LogDepth(0, level, msg)
}

func (log *logger) Debug(msg string, args ...any) {
	if log.Enabled(slog.DebugLevel) {
		log.Log(slog.DebugLevel, fmt.Sprintf(msg, args...))
	}
}

func (log *logger) Info(msg string, args ...any) {
	if log.Enabled(slog.InfoLevel) {
		log.Log(slog.InfoLevel, fmt.Sprintf(msg, args...))
	}
}

func (log *logger) Warn(err error) {
	if log.Enabled(slog.WarnLevel) {
		log.Log(slog.WarnLevel, err.Error())
	}
}

func (log *logger) Error(err error) {
	if log.Enabled(slog.ErrorLevel) {
		log.Log(slog.ErrorLevel, err.Error())
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
	if !logc.hasValuer && hasValuer(kvs...) {
		logc.hasValuer = true
	}
	logc.kvs = append(logc.kvs, kvs...)

	return logc
}

func (log *logger) Start() Logger {
	logc := log.copy()
	logc.timer = time.Now()

	return logc
}

func (log *logger) Stop() {
	cost := time.Now().Sub(log.timer).Milliseconds()

	log.With(
		"cost", fmt.Sprintf("%dms", cost),
		"caller", CallerFile(5, false),
	).Info("time-cost")
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
		slog:      log.slog,
		level:     log.level,
		hasValuer: log.hasValuer,
		kvs:       log.kvs,
	}
}
