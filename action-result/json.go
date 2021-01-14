package actionresult

import "github.com/xm-chentl/go-mvc"

// JSON 响应JSON
func JSON(data interface{}) mvc.IActionResult {
	return base{
		Err:  0,
		Data: data,
	}
}
