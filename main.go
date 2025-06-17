package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/arsolitt/slog-example/logger"
)

type Address struct {
	Host string
	Port int
}
type Req struct {
	Address   Address
	UserAgent string
	Path      string
}

func main() {
	// pass configuration here
	cfg := logger.NewLoggerConfig(slog.LevelDebug, true, true, true)
	logger.InitLogging(cfg)
	ctx := context.Background()

	reqID := "123121"
	// put something in the context
	ctx = logger.WithLogValue(ctx, logger.RequestIDField, reqID)
	slog.InfoContext(ctx, "New request")

	request := Req{
		Address: Address{
			Host: "localhost",
			Port: 8080,
		},
		UserAgent: "Mozilla/5.0",
		Path:      "/home",
	}

	// log debug
	slog.DebugContext(ctx, "Debug message before level changed")
	// we've done something and no longer need debug in this context. DON'T UNDERSTAND, BLOCKING
	ctx = logger.WithLogLevel(ctx, slog.LevelInfo)
	// no longer logging debug
	slog.DebugContext(ctx, "Debug message after level changed")

	// we can put a structure here and get a cool JSON that is just as easy to parse
	ctx = logger.WithLogValue(ctx, logger.RequestObject, request)
	slog.InfoContext(ctx, "Processing request")

	userId := "42"
	// use helper for field to avoid passing the name and get typing
	ctx = logger.WithLogUserID(ctx, userId)
	slog.InfoContext(ctx, "Processing user")

	instanceId := "228"
	ctx = logger.WithLogValue(ctx, logger.InstanceIDField, instanceId)
	slog.InfoContext(ctx, "Processing instance")

	// get an error
	err := errors.New("some error")
	// wrap error as desired
	err = fmt.Errorf("error wrapping: %w", err)
	// wrap the error again to put context in it. This can be done once at the place where the error occurred
	err = logger.CtxToError(ctx, err)
	// can wrap again
	err = fmt.Errorf("another error wrapping: %w", err)
	// log at the top level, get all the info
	slog.ErrorContext(logger.CtxFromError(ctx, err), err.Error())

	slog.InfoContext(ctx, "Done")
}
