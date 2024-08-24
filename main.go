package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/arsolitt/slog-example/logger"
)

func main() {
	// сюда прокидываем конфигурацию
	logger.InitLogging()
	ctx := context.Background()

	reqID := "123121"
	// запихиваем что-то в контекст
	ctx = logger.WithLogValue(ctx, logger.RequestIDField, reqID)
	slog.InfoContext(ctx, "New request")

	// логируем дебаг
	slog.DebugContext(ctx, "Debug message before level changed")
	// что-то навертели и больше нам дебаг в этом контексте не нужен. НЕ ПОНИМАЮ, БЛОКИРУЮ
	ctx = logger.WithLogLevel(ctx, slog.LevelInfo)
	// больше не логируем дебаг
	slog.DebugContext(ctx, "Debug message after level changed")

	userId := "42"
	// используем хелпер для поля, чтобы не прокидывать название и получить типизацию
	ctx = logger.WithLogUserID(ctx, userId)
	slog.InfoContext(ctx, "Processing user")

	instanceId := "228"
	ctx = logger.WithLogValue(ctx, logger.InstanceIDField, instanceId)
	slog.InfoContext(ctx, "Processing instance")

	// получаем ошибку
	err := errors.New("some error")
	// врапим ошибку по желанию
	err = fmt.Errorf("error wrapping: %w", err)
	// ещё раз врапим ошибку, чтобы положить в неё контекст. Можно сделать это один раз в том месте, где ошибка произошла
	err = logger.WrapError(ctx, err)
	// можно ещё раз заврапить
	err = fmt.Errorf("another error wrapping: %w", err)
	// логируем на самом верхнем уровне, получаем всю инфу
	slog.ErrorContext(logger.ErrorCtx(ctx, err), err.Error())

	slog.InfoContext(ctx, "Done")
}