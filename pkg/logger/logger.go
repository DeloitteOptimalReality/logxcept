package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ORLogger struct {
	zap.Logger
}

func NewLoggerWithTrace(trace string) *zap.Logger {
	logger := NewLogger()
	l := logger.With(zapcore.Field{Key: "requestTrace", Type: zapcore.StringType, String: trace})
	return l
}
func NewLogger() *zap.Logger {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
	return logger
}
