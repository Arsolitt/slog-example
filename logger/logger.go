package logger

import (
	"context"
	"log/slog"
	"os"
)

// сюда прокидываем конфигурацию
func InitLogging() {
	handler := slog.Handler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
	}))
	handler = NewMiddleware(handler)
	slog.SetDefault(slog.New(handler))
}

// сюда закидываем всякое
func WithLogValue(ctx context.Context, entryKey string, value any) context.Context {
	if c, ok := ctx.Value(dataKey).(logData); ok {
		c[entryKey] = value
		return context.WithValue(ctx, dataKey, c)
	}
	return context.WithValue(ctx, dataKey, logData{entryKey: value})
}

// опционально для конкретного поля
func WithLogUserID(ctx context.Context, userID string) context.Context {
	if c, ok := ctx.Value(dataKey).(logData); ok {
		c[UserIDField] = userID
		return context.WithValue(ctx, dataKey, c)
	}
	return context.WithValue(ctx, dataKey, logData{UserIDField: userID})
}

// можно прямо в рантайме переключать уровень логирования в разных частях приложения
func WithLogLevel(ctx context.Context, value slog.Level) context.Context {
	return context.WithValue(ctx, levelKey, value)
}
