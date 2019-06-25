package support

/**
* @author zhangsheng
* @date 2019/6/13
 */

type usResponse struct {
	Content interface{} `json:"content"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func Response(content interface{}, code int, message string) *usResponse {
	return &usResponse{
		Content: content,
		Code:    code,
		Message: message,
	}
}

func Success(content interface{}) *usResponse {
	return &usResponse{
		Content: content,
		Code:    1,
		Message: "Success",
	}
}

func Fail(err error) *usResponse {
	return &usResponse{
		Content: nil,
		Code:    -1,
		Message: err.Error(),
	}
}

func Error(mes string) *usResponse {
	return &usResponse{
		Content: nil,
		Code:    -1,
		Message: mes,
	}
}
