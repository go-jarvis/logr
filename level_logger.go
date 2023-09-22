package logr

import (
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

type levelLogger struct {
	logger *slog.Logger
	level  Level
}

// clone return a copy of levelLogger
func (l *levelLogger) clone() *levelLogger {
	c := *l
	return &c
}

// enabled will verify the level
func (l *levelLogger) enabled(lvl Level) bool {
	return l.level <= lvl
}

// output message
func (l *levelLogger) output(slevel slog.Level, msg string, args ...any) {
	l.logger.Log(nil, slevel, fmt.Sprintf(msg, args...))
}

func (l *levelLogger) Debug(msg string, args ...any) {
	if !l.enabled(DebugLevel) {
		return
	}

	l.output(slog.LevelDebug, msg, args...)
}

func (l *levelLogger) Info(msg string, args ...any) {
	if !l.enabled(InfoLevel) {
		return
	}

	l.output(slog.LevelInfo, msg, args...)
}

func (l *levelLogger) Warn(msg string, args ...any) {
	if !l.enabled(WarnLevel) {
		return
	}

	l.output(slog.LevelWarn, msg, args...)
}

func (l *levelLogger) Error(msg string, args ...any) {
	if !l.enabled(ErrorLevel) {
		return
	}

	l.output(slog.LevelError, msg, args...)
}

// Fatal output error log and exit.
func (l *levelLogger) Fatal(msg string, args ...any) {
	if !l.enabled(FatalLevel) {
		return
	}

	// args = append(args, "fatal", "true")

	l = l.with("fatal", "true")
	l.output(slog.LevelError, msg, args...)
	os.Exit(1)
}

// With k-v paris in log msg and return a new Logger
func (l *levelLogger) With(kvs ...any) Logger {

	return l.with(kvs...)
}

func (l *levelLogger) with(kvs ...any) *levelLogger {

	// clone
	cc := l.clone()

	// get new slog.Logger
	logcc := l.logger.With(kvs...)

	cc.logger = logcc
	return cc
}
