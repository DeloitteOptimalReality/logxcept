package middleware

import (
	"testing"
)

func TestMiddleware(t *testing.T) {
	mw := GinZapMiddlewareWithTrace()

	if mw == nil {
		t.Fail()
	}

	mw = GinZapMiddleware()
	if mw == nil {
		t.Fail()
	}
}
