package pojo

// 定义 Result 结构体
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 定义错误码
const (
	// 成功
	CODE_SUCCESS int = 1

	//失败
	CODE_ERROR int = -1

	//自定义...
)

// 设置错误码
func (res *Result) SetCode(code int) *Result {
	res.Code = code
	return res
}

// 设置错误信息
func (res *Result) SetMessage(msg string) *Result {
	res.Message = msg
	return res
}

// 设置返回数据
func (res *Result) SetData(data interface{}) *Result {
	res.Data = data
	return res
}

func (res *Result) Success(successMessage string, data interface{}) *Result {
	res.Message = successMessage
	res.Data = data
	res.Code = CODE_SUCCESS
	return res
}

func (res *Result) Failed(failedMessage string) *Result {
	res.Message = failedMessage
	res.Code = CODE_ERROR
	return res
}
