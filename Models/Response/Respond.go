package Response

func Respond(status int, message string, data interface{}) (res BaseResponse) {
	res = BaseResponse{
		Code:    status,
		Message: message,
		Data:    data,
	}
	return
}
