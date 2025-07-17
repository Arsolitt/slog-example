package logger

import (
	"context"
	"log/slog"
)

type ContextMiddleware struct {
	next slog.Handler
}

func NewContextMiddleware(next slog.Handler) *ContextMiddleware {
	return &ContextMiddleware{next: next}
}

func (h *ContextMiddleware) Enabled(ctx context.Context, rec slog.Level) bool {
	if level, ok := ctx.Value(levelKey).(slog.Level); ok {
		return rec >= level
	}
	return h.next.Enabled(ctx, rec)
}

func (h *ContextMiddleware) Handle(ctx context.Context, rec slog.Record) error {
	if c, ok := ctx.Value(dataKey).(*logData); ok {
		c.mu.RLock()
		defer c.mu.RUnlock()
		for k, v := range c.data {
			rec.Add(k, v)
		}
	}
	return h.next.Handle(ctx, rec)
}

func (h *ContextMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextMiddleware{next: h.next.WithAttrs(attrs)}
}

func (h *ContextMiddleware) WithGroup(name string) slog.Handler {
	return &ContextMiddleware{next: h.next.WithGroup(name)}
}
