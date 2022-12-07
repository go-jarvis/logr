package logr

import (
	"context"
	"os"
	"time"

	"golang.org/x/exp/slog"
)

var _ Logger = &logger{}

type logger struct {
	slog *slog.Logger

	timer     time.Time
	args      []any
	hasValuer bool
}

func Default() Logger {
	return &logger{
		slog: slog.Default(),
	}
}

func DefaultJSON() Logger {
	return &logger{
		slog: slog.New(slog.NewJSONHandler(os.Stdout)),
	}
}

func NewWithLogger(log *slog.Logger) Logger {
	return &logger{
		slog: log,
	}
}

func (log *logger) Debug(msg string, args ...any) {
	args = append(args, log.bindValuer()...)
	log.slog.Debug(msg, args...)
}

func (log *logger) Info(msg string, args ...any) {
	args = append(args, log.bindValuer()...)
	log.slog.Info(msg, args...)
}

func (log *logger) Warn(err error) {
	log.slog.Warn(err.Error(), log.bindValuer()...)
}

func (log *logger) Error(err error) {
	log.slog.Error(err.Error(), err, log.bindValuer()...)
}

// With params
func (log *logger) With(args ...any) Logger {
	if log.args == nil {
		log.args = make([]any, 0)
	}

	if hasValuer(args) {
		log.hasValuer = true
	}
	log.args = append(log.args, args...)
	return log
}

// Start 创建 log 副本， 并开启计时器
func (log *logger) Start(ctx context.Context, args ...any) (context.Context, Logger) {

	logc := log.copy().With(args...).(*logger)
	logc.timer = time.Now()

	return ctx, logc
}

func (log *logger) Stop() {

	cost := time.Now().Sub(log.timer).Milliseconds() / 1e3

	log.Info("time cost (second)",
		"cost", cost,
		"func", CallerFunc(2)(),
	)
}

// bindValuer
func (log *logger) bindValuer() []any {
	args := append([]any{}, log.args...)
	if log.hasValuer {
		return bindValuer(args)
	}
	return args
}

func (log *logger) copy() *logger {
	logc := &logger{
		slog:      log.slog,
		args:      log.args,
		hasValuer: log.hasValuer,
	}

	return logc
}
