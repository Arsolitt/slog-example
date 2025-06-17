package logger

import (
	"context"
	"log/slog"
	"os"
)

type loggerConfig struct {
	level       slog.Level
	isPretty    bool
	withContext bool
	withSources bool
}

// NewLoggerConfig create new logger configuration
func NewLoggerConfig(level slog.Level, isPretty bool, withContext bool, withSources bool) *loggerConfig {
	return &loggerConfig{
		level:       level,
		isPretty:    isPretty,
		withContext: withContext,
		withSources: withSources,
	}
}

// InitLogging initialize logger with configuration
func InitLogging(cfg *loggerConfig) {
	var handler slog.Handler
	if cfg.isPretty {
		opts := PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level:     cfg.level,
				AddSource: cfg.withSources,
			},
		}
		handler = opts.NewPrettyHandler(cfg.withContext, os.Stdout)
	} else {
		handler = slog.Handler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     cfg.level,
			AddSource: cfg.withSources,
		}))
	}

	if cfg.withContext {
		handler = NewContextMiddleware(handler)
	}
	slog.SetDefault(slog.New(handler))
}

// WithLogValue put anything here
func WithLogValue(ctx context.Context, entryKey string, value any) context.Context {
	if c, ok := ctx.Value(dataKey).(logData); ok {
		c[entryKey] = value
		return context.WithValue(ctx, dataKey, c)
	}
	return context.WithValue(ctx, dataKey, logData{entryKey: value})
}

// WithLogUserID optional for specific field
func WithLogUserID(ctx context.Context, userID string) context.Context {
	if c, ok := ctx.Value(dataKey).(logData); ok {
		c[UserIDField] = userID
		return context.WithValue(ctx, dataKey, c)
	}
	return context.WithValue(ctx, dataKey, logData{UserIDField: userID})
}

// WithLogLevel change log level in runtime
func WithLogLevel(ctx context.Context, value slog.Level) context.Context {
	return context.WithValue(ctx, levelKey, value)
}
