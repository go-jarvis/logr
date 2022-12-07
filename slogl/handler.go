package slogl

import "golang.org/x/exp/slog"

var _ slog.Handler = &LevelHandler{}
var _ slog.Leveler = &LevelHandler{}

type LevelHandler struct {
	level   slog.Level
	handler slog.Handler
}

func newLevelHanlder(l slog.Level, h slog.Handler) *LevelHandler {
	return &LevelHandler{
		level:   l,
		handler: h,
	}
}

func (lh *LevelHandler) Enabled(level slog.Level) bool {
	return lh.level > level
}

func (lh *LevelHandler) Handle(r slog.Record) error {
	return lh.handler.Handle(r)
}

func (h *LevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return newLevelHanlder(h.level, h.handler.WithAttrs(attrs))
}

func (h *LevelHandler) WithGroup(name string) slog.Handler {
	return newLevelHanlder(h.level, h.handler.WithGroup(name))
}

func (h *LevelHandler) Level() slog.Level {
	return h.level
}
