package logger

import (
	"context"
	"errors"
)

type errorWithLogCtx struct {
	next error
	data logData
}

func (e *errorWithLogCtx) Error() string {
	return e.next.Error()
}

func WrapError(ctx context.Context, err error) error {
	data := logData{}
	if d, ok := ctx.Value(dataKey).(logData); ok {
		data = d
	}
	return &errorWithLogCtx{
		next: err,
		data: data,
	}
}

func ErrorCtx(ctx context.Context, err error) context.Context {
	var e *errorWithLogCtx
	if errors.As(err, &e) {
		return context.WithValue(ctx, dataKey, e.data)
	}
	return ctx
}
