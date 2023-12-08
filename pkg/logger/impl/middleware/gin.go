package middleware

import (
	"context"
	"github.com/DeloitteOptimalReality/logxcept/internal/trace"
	"github.com/DeloitteOptimalReality/logxcept/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	CtxLoggerKey      = "orLogger"
	RequestLevelTrace = "requestTraceID"
)

func ZapMiddlewareWithTrace() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := trace.RandStringBytes(16)
		zapLog := logger.NewLoggerWithTrace(t)
		ctx := context.WithValue(c.Request.Context(), CtxLoggerKey, zapLog)
		ctx = context.WithValue(ctx, RequestLevelTrace, t)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		defer zapLog.Sync()
	}
}

func ZapMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		zapLog := logger.NewLogger()
		ctx := context.WithValue(c.Request.Context(), CtxLoggerKey, zapLog)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		defer zapLog.Sync()
	}
}

func RetrieveLogger(ctx context.Context) *zap.Logger {
	l := ctx.Value(CtxLoggerKey)
	if l == nil {
		return nil
	}

	zl := l.(*zap.Logger)
	return zl
}
