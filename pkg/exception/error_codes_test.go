package exception

import (
	"github.com/DeloitteOptimalReality/logxcept/pkg/exception/code"
	"testing"
)

func TestErrorCodes(t *testing.T) {
	if code.E1001.Code() != "E1001" || code.E1001.Msg() != "The system encountered an unexpected error. Please contact your admin." {
		t.Fail()
	}
	if code.E1002.Code() != "E1002" || code.E1002.Msg() != "Unable to communicate with server. Please try again." {
		t.Fail()
	}
	if code.E1003.Code() != "E1003" || code.E1003.Msg() != "Unable to communicate with database. Please try again." {
		t.Fail()
	}
	if code.E1004.Code() != "E1004" || code.E1004.Msg() != "Unable to perform action in database. Please try again." {
		t.Fail()
	}
	if code.E1005.Code() != "E1005" || code.E1005.Msg() != "Request timed out. Please try again." {
		t.Fail()
	}
	if code.E1006.Code() != "E1006" || code.E1006.Msg() != "Invalid argument. Please ensure the argument is correct." {
		t.Fail()
	}
	if code.E1007.Code() != "E1007" || code.E1007.Msg() != "The request was directed to a server that is unable to respond. Please ensure the request parameters are correct." {
		t.Fail()
	}
	if code.E1008.Code() != "E1008" || code.E1008.Msg() != "Too many requests have been made recently. Please try again later." {
		t.Fail()
	}
	if code.E1009.Code() != "E1009" || code.E1009.Msg() != "A dependency has failed. Please try again." {
		t.Fail()
	}
	if code.E2000.Code() != "E2000" || code.E2000.Msg() != "A syntax error has occurred. Please ensure the syntax is correct." {
		t.Fail()
	}
	if code.E2001.Code() != "E2001" || code.E2001.Msg() != "A configuration error has occurred. Please ensure the configuration is correct." {
		t.Fail()
	}
	if code.E2002.Code() != "E2002" || code.E2002.Msg() != "Invalid request received. Please ensure the request is correct." {
		t.Fail()
	}
	if code.E2003.Code() != "E2003" || code.E2003.Msg() != "Insufficient permissions for user to perform this action." {
		t.Fail()
	}
}
