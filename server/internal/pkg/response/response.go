package response

const (
	success = 200
	failure = 500
)

type ServerResponse struct {
	ResponseCode
	Data interface{} `json:"data"`
}

func CreateBySuccess() ServerResponse {
	return ServerResponse{SuccessCode, nil}
}

func CreateBySuccessMessage(msg string) ServerResponse {
	return ServerResponse{ResponseCode{success, msg}, nil}
}

func CreateBySuccessCodeMessage(code int, msg string) ServerResponse {
	return ServerResponse{ResponseCode{code, msg}, nil}
}

func CreateBySuccessData(data interface{}) ServerResponse {
	return ServerResponse{SuccessCode, data}
}

func CreateByError() ServerResponse {
	return ServerResponse{FailureCode, nil}
}

func CreateByErrorMessage(err error) ServerResponse {
	return ServerResponse{ResponseCode{failure, err.Error()}, nil}
}

func CreateByErrorCodeMessage(rc ResponseCode) ServerResponse {
	return ServerResponse{rc, nil}
}
