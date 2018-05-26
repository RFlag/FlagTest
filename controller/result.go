package controller

var (
	ResultOk                       = map[string]int{"code": 1}
	ResultParamError               = ResultError("参数错误")
	ResultHttpResuestCreateFailed  = ResultError("创建HTTP请求失败")
	ResultHttpResuestExecuteFailed = ResultError("创建HTTP请求失败")
)

type resultCodeError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

var resultCode = 9

func ResultError(e string) resultCodeError {
	resultCode++
	return resultCodeError{
		Code:  resultCode,
		Error: e,
	}
}

func ResultData(d interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Code": 1,
		"Data": d,
	}
}
