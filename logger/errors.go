package logger

import (
	"context"
	"errors"
	"sync"
)

type CtxError struct {
	next error
	data map[string]any
}

func (e *CtxError) Error() string {
	return e.next.Error()
}

func CtxToError(ctx context.Context, err error) error {
	data := map[string]any{}
	if d, ok := ctx.Value(dataKey).(*logData); ok {
		d.mu.RLock()
		defer d.mu.RUnlock()
		data = d.data
	}
	return &CtxError{
		next: err,
		data: data,
	}
}

func CtxFromError(ctx context.Context, err error) context.Context {
	var e *CtxError
	if errors.As(err, &e) {
		return context.WithValue(ctx, dataKey, &logData{data: e.data, mu: sync.RWMutex{}})
	}
	return ctx
}
