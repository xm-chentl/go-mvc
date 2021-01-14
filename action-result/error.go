package actionresult

import (
	"fmt"

	"github.com/xm-chentl/go-mvc"
)

// Error 错误
func Error(code int, format string, args ...interface{}) mvc.IActionResult {
	return &base{
		Err:  code,
		Data: fmt.Sprintf(format, args...),
	}
}
