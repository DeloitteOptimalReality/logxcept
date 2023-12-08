package test_logger

import (
	"github.com/DeloitteOptimalReality/logxcept/pkg/logger/impl/middleware"
	"testing"
)

func TestMiddleware(t *testing.T) {
	mw := middleware.ZapMiddlewareWithTrace()

	if mw == nil {
		t.Fail()
	}

	mw = middleware.ZapMiddleware()
	if mw == nil {
		t.Fail()
	}
}
