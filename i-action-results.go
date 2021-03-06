package mvc

// IActionResult 响应接口
type IActionResult interface {
	Execute() interface{}
}
