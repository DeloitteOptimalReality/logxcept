package code

// code for exceptions

type ErrorCode struct {
	code    string
	message string
}

var (
	E1001 = ErrorCode{code: "E1001", message: "The system encountered an unexpected error. Please contact your admin."}
	E1002 = ErrorCode{code: "E1002", message: "Unable to communicate with server. Please try again."}
	E1003 = ErrorCode{code: "E1003", message: "Unable to communicate with database. Please try again."}
	E1004 = ErrorCode{code: "E1004", message: "Unable to perform action in database. Please try again."}
	E1005 = ErrorCode{code: "E1005", message: "Request timed out. Please try again."}
	E1006 = ErrorCode{code: "E1006", message: "Invalid argument. Please ensure the argument is correct."}
	E1007 = ErrorCode{code: "E1007", message: "The request was directed to a server that is unable to respond. Please ensure the request parameters are correct."}
	E1008 = ErrorCode{code: "E1008", message: "Too many requests have been made recently. Please try again later."}
	E1009 = ErrorCode{code: "E1009", message: "A dependency has failed. Please try again."}
	E2000 = ErrorCode{code: "E2000", message: "A syntax error has occurred. Please ensure the syntax is correct."}
	E2001 = ErrorCode{code: "E2001", message: "A configuration error has occurred. Please ensure the configuration is correct."}
	E2002 = ErrorCode{code: "E2002", message: "Invalid request received. Please ensure the request is correct."}
	E2003 = ErrorCode{code: "E2003", message: "Insufficient permissions for user to perform this action."}
)

func (e *ErrorCode) Code() string {
	return e.code
}

func (e *ErrorCode) Msg() string {
	return e.message
}
