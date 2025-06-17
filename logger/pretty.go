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
	// TODO: add options
	// opts PrettyHandlerOptions
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

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error

	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if h.withContext {
		return &PrettyHandler{
			Handler:     h.Handler,
			l:           h.l,
			attrs:       attrs,
			withContext: h.withContext,
		}
	}
	return &PrettyHandler{
		Handler:     h.Handler,
		l:           h.l,
		withContext: h.withContext,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// TODO: implement
	return &PrettyHandler{
		Handler:     h.Handler.WithGroup(name),
		l:           h.l,
		withContext: h.withContext,
	}
}
