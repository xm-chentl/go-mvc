package actionresult

import (
	"fmt"

	"github.com/xm-chentl/go-mvc"
)

type baseResult struct {
	Code int
	Msg  string
	Data interface{}
}

func (r baseResult) Execute() interface{} {
	return map[string]interface{}{
		"code": r.Code,
		"msg":  r.Msg,
		"data": r.Data,
	}
}

// Alert 提示级别的响应
func Alert(code int, msg string) mvc.IActionResult {
	return &baseResult{
		Code: code,
		Msg:  msg,
	}
}

// Alertf 提示级别的响应
func Alertf(code int, format string, args ...interface{}) mvc.IActionResult {
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	return &baseResult{
		Code: code,
		Msg:  msg,
	}
}

// JSON 响应json数据
func JSON(data interface{}) mvc.IActionResult {
	return &baseResult{
		Data: data,
	}
}
