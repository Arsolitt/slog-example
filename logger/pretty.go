package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l           *log.Logger
	attrs       []slog.Attr
	withContext bool
}

func (opts PrettyHandlerOptions) NewPrettyHandler(
	withContext bool,
	out io.Writer,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler:     slog.NewJSONHandler(out, opts.SlogOpts),
		l:           log.New(out, "", 0),
		withContext: withContext,
		attrs:       make([]slog.Attr, 0),
	}

	return h
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs()+len(h.attrs))

	// Добавляем attrs из handler'а
	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	// Добавляем attrs из record'а
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	var fieldsStr string
	if len(fields) > 0 {
		b, err := json.MarshalIndent(fields, "", "  ")
		if err != nil {
			fieldsStr = color.RedString("error marshaling fields: %v", err)
		} else {
			fieldsStr = color.WhiteString(string(b))
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, fieldsStr)
	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if !h.withContext {
		return h
	}

	// Объединяем существующие attrs с новыми
	newAttrs := make([]slog.Attr, 0, len(h.attrs)+len(attrs))
	newAttrs = append(newAttrs, h.attrs...)
	newAttrs = append(newAttrs, attrs...)

	return &PrettyHandler{
		Handler:     h.Handler,
		l:           h.l,
		attrs:       newAttrs,
		withContext: h.withContext,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	return &PrettyHandler{
		Handler:     h.Handler.WithGroup(name),
		l:           h.l,
		attrs:       h.attrs,
		withContext: h.withContext,
	}
}
