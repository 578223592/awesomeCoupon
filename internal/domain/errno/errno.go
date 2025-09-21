package errno

// CommonErrorCode 定义通用错误码
type CommonErrorCode int

const (
	Success             CommonErrorCode = 0
	UnknownError        CommonErrorCode = 100
	InvalidParam        CommonErrorCode = 110
	NotFound            CommonErrorCode = 404
	InternalServerError CommonErrorCode = 500
)

// String 返回错误码对应的错误信息
func (e CommonErrorCode) String() string {
	switch e {
	case Success:
		return "成功"
	case InvalidParam:
		return "参数错误"
	case NotFound:
		return "资源不存在"
	case InternalServerError:
		return "服务器内部错误"
	default:
		return "未知错误"
	}
}

// ErrorCode 定义错误码接口
type ErrorCode interface {
	Code() int
	Message() string
}

// CommonError 通用错误结构体
type CommonError struct {
	code    CommonErrorCode
	message string
}

func (e *CommonError) Code() int {
	if e == nil {
		return int(Success)
	}
	return int(e.code)
}

func (e *CommonError) Message() string {
	if e == nil {
		return Success.String()
	}
	if e.message != "" {
		return e.message
	}
	return e.code.String()
}

func NewCommonError(code CommonErrorCode, message string) *CommonError {
	err := &CommonError{code: code}
	if len(message) > 0 {
		err.message = message
	}
	return err
}
