package logger

import (
	"context"
	"log/slog"
)

type contextMiddleware struct {
	next slog.Handler
}

func NewContextMiddleware(next slog.Handler) *contextMiddleware {
	return &contextMiddleware{next: next}
}

func (h *contextMiddleware) Enabled(ctx context.Context, rec slog.Level) bool {
	if level, ok := ctx.Value(levelKey).(slog.Level); ok {
		return rec >= level
	}
	return h.next.Enabled(ctx, rec)
}

func (h *contextMiddleware) Handle(ctx context.Context, rec slog.Record) error {
	if c, ok := ctx.Value(dataKey).(logData); ok {
		for k, v := range c {
			rec.Add(k, v)
		}
	}
	return h.next.Handle(ctx, rec)
}

func (h *contextMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &contextMiddleware{next: h.next.WithAttrs(attrs)}
}

func (h *contextMiddleware) WithGroup(name string) slog.Handler {
	return &contextMiddleware{next: h.next.WithGroup(name)}
}
