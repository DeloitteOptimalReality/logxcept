package test_logger

import (
	"github.com/DeloitteOptimalReality/logxcept/pkg/logger"
	"testing"
)

func TestLoggerInit(t *testing.T) {
	zl := logger.NewLogger()
	if zl == nil {
		t.Fail()
	}

	zlt, tr := logger.NewLoggerWithTrace()

	if len(tr) != 16 {
		t.Fail()
	}
	if zlt == nil {
		t.Fail()
	}

}
