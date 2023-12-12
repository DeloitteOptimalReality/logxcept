package exception

import (
	"errors"
	"github.com/DeloitteOptimalReality/logxcept/pkg/exception/code"
	"github.com/DeloitteOptimalReality/logxcept/pkg/exception/impl"
	"testing"
)

func TestBaseException(t *testing.T) {
	be := NewBaseException(nil, "", nil, nil, nil)

	if be.Err() != nil {
		t.Fail()
	}

	e := errors.New("this is an example error that should")

	be = NewBaseException(e, "An error occurred in testing", nil, nil, nil)

	if be.Error() == be.Msg() {
		t.Fail()
	}

	return
}

func TestDatabaseException(t *testing.T) {
	dbe := impl.NewDatabaseException(nil, "test", nil, nil)

	// test the error is actually empty
	if dbe.Err() != nil {
		t.Fail()
	}
	c := dbe.Code()
	if c.Code() != code.E1003.Code() {
		t.Fail()

	}

	dbe = impl.NewDatabaseActionException(nil, "test", nil, nil)
	c = dbe.Code()
	if c.Code() != code.E1004.Code() {
		t.Fail()

	}

	err := errors.New("database error that has occurred")

	dbe = impl.NewDatabaseException(err, "something happened in the database", nil, nil)
	var dbePointer *impl.DatabaseException
	if !errors.As(dbe, &dbePointer) {
		t.Fail()
		//error should appear as a database error
	}

	var bePointer *BaseException

	if errors.As(dbe, &bePointer) {
		t.Fail()
		//error should not appear as a baseException error
	}
	if !errors.As(dbe.BaseException, &bePointer) {
		t.Fail()
		//error at the core should be a baseException error
	}

}
