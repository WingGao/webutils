package flogger

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gookit/goutil/errorx"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	My         *zap.Logger
	MyNoCaller *zap.Logger
)

func InitLogger() {
	c := zap.NewDevelopmentEncoderConfig()
	c.EncodeLevel = zapcore.CapitalColorLevelEncoder
	base := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(c),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))
	My = base.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	initLogger2()
}

func initLogger2() {
	c := zap.NewDevelopmentEncoderConfig()
	c.EncodeLevel = zapcore.CapitalColorLevelEncoder
	base := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(c),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	), zap.Hooks(func(entry zapcore.Entry) error {
		return nil
	}))
	MyNoCaller = base
}

func WithFiber(c fiber.Ctx, lg ...*zap.Logger) *zap.Logger {
	mlog := My
	if len(lg) > 0 {
		mlog = lg[0]
	}
	return mlog.With(zap.String("request_id", c.GetRespHeader("X-Request-ID")))
}

func WithFiberError(c fiber.Ctx, e error) {
	if ex, ok := e.(*errorx.ErrorX); ok {
		WithFiber(c, MyNoCaller).Error(ex.Cause().Error() + ex.StackString())
	}
	WithFiber(c).Error(e.Error())
}
