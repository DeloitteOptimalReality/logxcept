package impl

import (
	"context"
	"github.com/DeloitteOptimalReality/logxcept/pkg/exception"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type DatabaseException struct {
	*exception.BaseException
}

func NewDatabaseException(err error, msg string, path ast.Path, source string) *DatabaseException {
	be := exception.NewBaseException(err, msg, path, source)
	return &DatabaseException{be}
}

func (dbe *DatabaseException) Error() string {
	return dbe.BaseException.Error()
}

func (dbe *DatabaseException) Err() error {
	return dbe.BaseException.Err()
}

func (dbe *DatabaseException) Msg() string {
	return dbe.BaseException.Msg()
}

func (dbe *DatabaseException) Path() ast.Path {
	return dbe.BaseException.Path()
}

func (dbe *DatabaseException) Source() string {
	return dbe.BaseException.Error()
}

func (dbe *DatabaseException) Log(ctx context.Context) {
	dbe.BaseException.Log(ctx)
}

func (dbe *DatabaseException) GqlError() *gqlerror.Error {
	return dbe.BaseException.GqlError()
}

func (dbe *DatabaseException) GqlErrorWithTrace(ctx context.Context) *gqlerror.Error {
	return dbe.BaseException.GqlErrorWithTrace(ctx)
}
