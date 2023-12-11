package test_exception

import (
	"github.com/DeloitteOptimalReality/logxcept/pkg/exception"
	"testing"
)

func TestErrorCodes(t *testing.T) {
	if exception.E1001.Code() != "E1001" || exception.E1001.Msg() != "The system encountered an unexpected error. Please contact your admin." {
		t.Fail()
	}
	if exception.E1002.Code() != "E1002" || exception.E1002.Msg() != "Unable to communicate with server. Please try again." {
		t.Fail()
	}
	if exception.E1003.Code() != "E1003" || exception.E1003.Msg() != "Unable to communicate with database. Please try again." {
		t.Fail()
	}
	if exception.E1004.Code() != "E1004" || exception.E1004.Msg() != "Unable to perform action in database. Please try again." {
		t.Fail()
	}
	if exception.E1005.Code() != "E1005" || exception.E1005.Msg() != "Request timed out. Please try again." {
		t.Fail()
	}
	if exception.E1006.Code() != "E1006" || exception.E1006.Msg() != "Invalid argument. Please ensure the argument is correct." {
		t.Fail()
	}
	if exception.E1007.Code() != "E1007" || exception.E1007.Msg() != "The request was directed to a server that is unable to respond. Please ensure the request parameters are correct." {
		t.Fail()
	}
	if exception.E1008.Code() != "E1008" || exception.E1008.Msg() != "Too many requests have been made recently. Please try again later." {
		t.Fail()
	}
	if exception.E1009.Code() != "E1009" || exception.E1009.Msg() != "A dependency has failed. Please try again." {
		t.Fail()
	}
	if exception.E2000.Code() != "E2000" || exception.E2000.Msg() != "A syntax error has occurred. Please ensure the syntax is correct." {
		t.Fail()
	}
	if exception.E2001.Code() != "E2001" || exception.E2001.Msg() != "A configuration error has occurred. Please ensure the configuration is correct." {
		t.Fail()
	}
	if exception.E2002.Code() != "E2002" || exception.E2002.Msg() != "Invalid request received. Please ensure the request is correct." {
		t.Fail()
	}
	if exception.E2003.Code() != "E2003" || exception.E2003.Msg() != "Insufficient permissions for user to perform this action." {
		t.Fail()
	}
}
