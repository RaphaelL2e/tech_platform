package response

type ResponseCode struct {
	Code    int
	Message string
}

var SuccessCode = ResponseCode{200, "success"}
var FailureCode = ResponseCode{500, "failure"}
var ForbiddenCode = ResponseCode{401,"invalid token"}
var StatusForbiddenCode = ResponseCode{402,"status forbidden"}
var LoginErrCode = ResponseCode{4001,"login error: not found"}
var AdminErrCode = ResponseCode{4002,"login error: not is admin"}
var TokenGenErrCode = ResponseCode{4003,"token error: gen token failure"}

