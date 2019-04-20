package response

type Response struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
}

func SendError(status int, code error, message interface{}) *Response {
	return &Response{
		Status:  status,
		Code:    code.Error(),
		Message: message,
	}
}

func SendSuccess(status int, code error, message interface{}) *Response {
	return &Response{
		Status:  status,
		Code:    code.Error(),
		Message: message,
	}
}
