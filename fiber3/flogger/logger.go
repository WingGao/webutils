package flogger

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	My *zap.Logger
)

func InitLogger() {
	c := zap.NewDevelopmentEncoderConfig()
	c.EncodeLevel = zapcore.CapitalColorLevelEncoder
	My = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(c),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	), zap.AddCaller(), zap.AddCallerSkip(1))
}

func WithFiber(c fiber.Ctx) *zap.Logger {
	return My.With(zap.String("request_id", c.GetRespHeader("X-Request-ID")))
}

func WithFiberError(c fiber.Ctx, e error) {
	WithFiber(c).Error(e.Error())
}
