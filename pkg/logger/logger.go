package logger

import (
	"github.com/DeloitteOptimalReality/logxcept/internal/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ORLogger struct {
	zap.Logger
}

func NewLoggerWithTrace(level *string) (*zap.Logger, string) {
	logger := NewLogger(level)
	t := trace.RandStringBytes(16)
	l := logger.With(zapcore.Field{Key: "requestTrace", Type: zapcore.StringType, String: t})
	return l, t
}
func NewLogger(level *string) *zap.Logger {
	logger := getLogger(level)
	defer logger.Sync()
	return logger
}

func getLogger(level *string) *zap.Logger {
	logger := InfoLogger()
	if level != nil {
		switch l, _ := zapcore.ParseLevel(*level); l {
		case zap.DebugLevel:
			logger = DebugLogger()
			break
		case zap.InfoLevel:
			logger = InfoLogger()
			break
		case zap.WarnLevel:
			logger = WarnLogger()
			break
		default:
			logger = InfoLogger()
		}
	}
	return logger
}

func DebugLogger() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	logger, _ := cfg.Build()
	return logger
}

func InfoLogger() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	logger, _ := cfg.Build()
	return logger
}

func WarnLogger() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	logger, _ := cfg.Build()
	return logger
}
