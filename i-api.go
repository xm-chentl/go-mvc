package mvc

import "github.com/xm-chentl/go-mvc/scope"

// IApi 接口规范
type IApi interface {
	Code() string
	Scope() scope.Value
	Execute() IActionResult
}
