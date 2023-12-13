package exception

import (
	"context"
	"fmt"
	"github.com/DeloitteOptimalReality/logxcept/internal/trace"
	"github.com/DeloitteOptimalReality/logxcept/pkg/exception/code"
	"github.com/DeloitteOptimalReality/logxcept/pkg/logger/impl/middleware"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Exception interface {
	Error() string
	Err() error
	Msg() string
	Path() ast.Path
	Source() string
	Code() *code.ErrorCode
	Log(ctx context.Context)
	GqlError() *gqlerror.Error
	GqlErrorWithTrace(ctx context.Context) *gqlerror.Error
}

var traceIDField = "traceID"
var resolverTraceIDField = "resolverTraceID"

type BaseException struct {
	err     error
	msg     string
	path    ast.Path
	source  *string
	traceID string
	code    *code.ErrorCode
}

func NewBaseException(err error, msg string, path ast.Path, code *code.ErrorCode, source *string) *BaseException {
	return &BaseException{
		err:     err,
		msg:     msg,
		path:    path,
		source:  source,
		traceID: trace.RandStringBytes(16),
		code:    code,
	}
}

func (be *BaseException) Error() string {
	return be.err.Error()
}

func (be *BaseException) Err() error {
	return be.err
}

func (be *BaseException) Msg() string {
	return be.msg
}

func (be *BaseException) Path() ast.Path {
	return be.path
}

func (be *BaseException) Source() string {
	return *be.source
}
func (be *BaseException) Code() *code.ErrorCode {
	return be.code
}

func (be *BaseException) Log(ctx context.Context) {

	var logger *zap.Logger

	l := ctx.Value(middleware.CtxLoggerKey)
	if l == nil {
		return
	}
	logger = l.(*zap.Logger)

	resolverTrace := ctx.Value(middleware.RequestLevelTrace).(string)

	// Include trace ID in the log
	fields := []zapcore.Field{
		zap.String(traceIDField, be.traceID),
		zap.String(resolverTraceIDField, resolverTrace),
		zap.String(be.Code().Code(), be.Code().Msg()),
	}

	logger.With(fields...).Info(fmt.Sprintf("%s: %s", be.code.Code(), be.code.Msg()))
	if logger.Level().String() != "info" {
		logger.Error(fmt.Sprintf("%s: %s", be.Msg(), be.err.Error()), fields...)
	}
}

func (be *BaseException) GqlError() *gqlerror.Error {
	return &gqlerror.Error{
		Err:       be.err,
		Message:   be.msg,
		Path:      be.path,
		Locations: nil,
		Extensions: map[string]interface{}{
			be.code.Code(): be.code.Msg(),
			traceIDField:   be.traceID,
		},
		Rule: "",
	}
}

func (be *BaseException) GqlErrorWithTrace(ctx context.Context) *gqlerror.Error {

	return &gqlerror.Error{
		Err:       be.err,
		Message:   be.msg,
		Path:      be.path,
		Locations: nil,
		Extensions: map[string]interface{}{
			be.code.Code():       be.code.Msg(),
			traceIDField:         be.traceID,
			resolverTraceIDField: traceFromCtx(ctx),
		},
		Rule: "",
	}
}

func traceFromCtx(ctx context.Context) string {
	resolverTrace := ""

	r := ctx.Value(middleware.RequestLevelTrace)
	if r == nil {
		return resolverTrace
	}

	resolverTrace = r.(string)
	return resolverTrace
}
