package exception

import (
	"context"
	"fmt"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"math/rand"
)

type Exception interface {
	Error() string
	Err() error
	Msg() string
	Path() ast.Path
	Source() string
	Log(ctx context.Context)
	GqlError() *gqlerror.Error
	GqlErrorWithTrace(ctx context.Context) *gqlerror.Error
}

var traceIDField = "traceID"
var resolverTraceIDField = "resolverTraceID"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type BaseException struct {
	err     error
	msg     string
	path    ast.Path
	source  string
	traceID string
}

func NewBaseException(err error, msg string, path ast.Path, source string) *BaseException {
	return &BaseException{
		err:     err,
		msg:     msg,
		path:    path,
		source:  source,
		traceID: RandStringBytes(16),
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
	return be.source
}

func (be *BaseException) Log(ctx context.Context) {
	var logger *zap.Logger

	l := ctx.Value("logger")
	if l == nil {
		return
	}
	logger = l.(*zap.Logger)

	resolverTrace := ctx.Value("resolverTrace").(string)

	// Include trace ID in the log
	fields := []zapcore.Field{
		zap.String(traceIDField, be.traceID),
		zap.String("resolverTrace", resolverTrace),
	}

	logger.With(fields...).Info(fmt.Sprintf("%s: %s", be.Msg(), be.err.Error()))
}

func (be *BaseException) GqlError() *gqlerror.Error {
	return &gqlerror.Error{
		Err:       be.err,
		Message:   be.msg,
		Path:      be.path,
		Locations: nil,
		Extensions: map[string]interface{}{
			traceIDField: be.traceID,
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
			traceIDField:         be.traceID,
			resolverTraceIDField: traceFromCtx(ctx),
		},
		Rule: "",
	}
}

func traceFromCtx(ctx context.Context) string {
	resolverTrace := ""

	r := ctx.Value("resolverTrace")
	if r == nil {
		return resolverTrace
	}

	resolverTrace = r.(string)
	return resolverTrace
}
