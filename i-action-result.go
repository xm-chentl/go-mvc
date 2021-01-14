package mvc

// IActionResult 响应掊口
type IActionResult interface {
	Exec() interface{}
}
