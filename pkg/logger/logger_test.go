package logger

import (
	"testing"
)

func TestLoggerInit(t *testing.T) {
	zl := NewLogger()
	if zl == nil {
		t.Fail()
	}

	zlt, tr := NewLoggerWithTrace()

	if len(tr) != 16 {
		t.Fail()
	}
	if zlt == nil {
		t.Fail()
	}

}
