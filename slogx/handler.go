package slogx

import (
	"context"

	"golang.org/x/exp/slog"
)

var _ slog.Handler = &slogHandler{}

type slogHandler struct {
	h slog.Handler
}

func DefaultHandler() slog.Handler {
	return NewHandler(slog.Default().Handler())
}

func NewHandler(handler slog.Handler) slog.Handler {
	return &slogHandler{
		h: handler,
	}
}

func (h *slogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (h *slogHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.h.Handle(ctx, r)
}

func (h *slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewHandler(h.h.WithAttrs(attrs))
}

func (h *slogHandler) WithGroup(name string) slog.Handler {
	return NewHandler(h.h.WithGroup(name))
}
