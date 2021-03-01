package response

type ResponseCode struct {
	Code    int
	Message string
}

var SuccessCode = ResponseCode{200, "success"}
var FailureCode = ResponseCode{500, "failure"}
var ForbiddenCode = ResponseCode{401,"invalid token"}
var StatusForbiddenCode = ResponseCode{402,"status forbidden"}
