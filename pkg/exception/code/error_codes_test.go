package code

import (
	"testing"
)

func TestErrorCodes(t *testing.T) {
	if E1001.Code() != "E1001" || E1001.Msg() != "The system encountered an unexpected error. Please contact your admin." {
		t.Fail()
	}
	if E1002.Code() != "E1002" || E1002.Msg() != "Unable to communicate with server. Please try again." {
		t.Fail()
	}
	if E1003.Code() != "E1003" || E1003.Msg() != "Unable to communicate with database. Please try again." {
		t.Fail()
	}
	if E1004.Code() != "E1004" || E1004.Msg() != "Unable to perform action in database. Please try again." {
		t.Fail()
	}
	if E1005.Code() != "E1005" || E1005.Msg() != "Request timed out. Please try again." {
		t.Fail()
	}
	if E1006.Code() != "E1006" || E1006.Msg() != "Invalid argument. Please ensure the argument is correct." {
		t.Fail()
	}
	if E1007.Code() != "E1007" || E1007.Msg() != "The request was directed to a server that is unable to respond. Please ensure the request parameters are correct." {
		t.Fail()
	}
	if E1008.Code() != "E1008" || E1008.Msg() != "Too many requests have been made recently. Please try again later." {
		t.Fail()
	}
	if E1009.Code() != "E1009" || E1009.Msg() != "A dependency has failed. Please try again." {
		t.Fail()
	}
	if E2000.Code() != "E2000" || E2000.Msg() != "A syntax error has occurred. Please ensure the syntax is correct." {
		t.Fail()
	}
	if E2001.Code() != "E2001" || E2001.Msg() != "A configuration error has occurred. Please ensure the configuration is correct." {
		t.Fail()
	}
	if E2002.Code() != "E2002" || E2002.Msg() != "Invalid request received. Please ensure the request is correct." {
		t.Fail()
	}
	if E2003.Code() != "E2003" || E2003.Msg() != "Insufficient permissions for user to perform this action." {
		t.Fail()
	}
}
