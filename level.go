package logr

import (
	"strconv"
	"strings"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// func (Level).Level() Level
// func (Level).MarshalJSON() ([]byte, error)
// func (Level).String() string

func (l Level) Level() Level {
	return l
}
func (l Level) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, l.String()), nil
}

func (l Level) String() string {
	return LevelToText(l)
}

func LevelToText(level Level) string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	}

	return "info"
}

func LevelFromText(level string) Level {
	switch strings.ToLower(level) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	}

	return InfoLevel
}
