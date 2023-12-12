package impl

import (
	"errors"
	"github.com/DeloitteOptimalReality/logxcept/pkg/exception"
	"github.com/DeloitteOptimalReality/logxcept/pkg/exception/code"
	"testing"
)

func TestDatabaseException(t *testing.T) {
	dbe := NewDatabaseException(nil, "test", nil, nil)

	// test the error is actually empty
	if dbe.Err() != nil {
		t.Fail()
	}
	c := dbe.Code()
	if c.Code() != code.E1003.Code() {
		t.Fail()

	}

	dbe = NewDatabaseActionException(nil, "test", nil, nil)
	c = dbe.Code()
	if c.Code() != code.E1004.Code() {
		t.Fail()

	}

	err := errors.New("database error that has occurred")

	dbe = NewDatabaseException(err, "something happened in the database", nil, nil)
	var dbePointer *DatabaseException
	if !errors.As(dbe, &dbePointer) {
		t.Fail()
		//error should appear as a database error
	}

	var bePointer *exception.BaseException

	if errors.As(dbe, &bePointer) {
		t.Fail()
		//error should not appear as a baseException error
	}
	if !errors.As(dbe.BaseException, &bePointer) {
		t.Fail()
		//error at the core should be a baseException error
	}

}
