package exception

import (
	"errors"
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
