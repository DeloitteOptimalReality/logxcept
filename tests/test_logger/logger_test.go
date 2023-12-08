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

	zlt := logger.NewLoggerWithTrace()

	if zlt == nil {
		t.Fail()
	}

}
