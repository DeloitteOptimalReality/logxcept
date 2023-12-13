package logger

import (
	"testing"
)

func TestLoggerInit(t *testing.T) {
	level := "debug"
	zl := NewLogger(&level)
	if zl == nil {
		t.Fail()
	}

	if zl.Level().String() != level {
		t.Fail()
	}

	level = "warn"
	zlt, tr := NewLoggerWithTrace(&level)

	if len(tr) != 16 {
		t.Fail()
	}
	if zlt == nil {
		t.Fail()
	}

	if zlt.Level().String() != level {
		t.Fail()
	}

	level = "info"
	zlti, tr := NewLoggerWithTrace(&level)

	if zlti.Level().String() != level {
		t.Fail()
	}

	zltn, _ := NewLoggerWithTrace(nil)
	if zltn.Level().String() != level {
		t.Fail()
	}

}

func TestLoggerLevel(t *testing.T) {
	l := InfoLogger()
	if l.Level().String() != "info" {
		t.Fail()
	}
	l = DebugLogger()
	if l.Level().String() != "debug" {
		t.Fail()
	}
	l = WarnLogger()
	if l.Level().String() != "warn" {
		t.Fail()
	}
}
