package logger

import (
	"github.com/DeloitteOptimalReality/logxcept/internal/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ORLogger struct {
	zap.Logger
}

func NewLoggerWithTrace() *zap.Logger {
	logger := NewLogger()
	t := trace.RandStringBytes(16)
	l := logger.With(zapcore.Field{Key: "requestTrace", Type: zapcore.StringType, String: t})
	return l
}
func NewLogger() *zap.Logger {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
	return logger
}
