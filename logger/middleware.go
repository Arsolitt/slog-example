package logger

import (
	"context"
	"log/slog"
)

type Middlware struct {
	next slog.Handler
}

func NewMiddleware(next slog.Handler) *Middlware {
	return &Middlware{next: next}
}

func (h *Middlware) Enabled(ctx context.Context, rec slog.Level) bool {
	if level, ok := ctx.Value(levelKey).(slog.Level); ok {
		return rec >= level
	}
	return h.next.Enabled(ctx, rec)
}

func (h *Middlware) Handle(ctx context.Context, rec slog.Record) error {
	if c, ok := ctx.Value(dataKey).(logData); ok {
		for k, v := range c {
			rec.Add(k, v)
		}
	}
	return h.next.Handle(ctx, rec)
}

func (h *Middlware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Middlware{next: h.next.WithAttrs(attrs)}
}

func (h *Middlware) WithGroup(name string) slog.Handler {
	return &Middlware{next: h.next.WithGroup(name)}
}
