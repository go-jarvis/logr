package slogx

import "golang.org/x/exp/slog"

var _ slog.Handler = &slogDefaultHandler{}

type slogDefaultHandler struct {
	handler slog.Handler
}

func defaultHandler() slog.Handler {
	return newDefaultHandler(slog.Default().Handler())
}

func newDefaultHandler(handler slog.Handler) slog.Handler {
	return &slogDefaultHandler{
		handler: handler,
	}
}

func (h *slogDefaultHandler) Enabled(level slog.Level) bool {
	return true
}

func (h *slogDefaultHandler) Handle(r slog.Record) error {
	return h.handler.Handle(r)
}

func (h *slogDefaultHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return newDefaultHandler(h.handler.WithAttrs(attrs))
}

func (h *slogDefaultHandler) WithGroup(name string) slog.Handler {
	return newDefaultHandler(h.handler.WithGroup(name))
}
