package handler

import (
	"github.com/xm-chentl/go-mvc"
)

// Default 默认配置
func Default() mvc.IHandler {
	handler := new(Handle)
	handler.Next(
		new(Service),
	).Next(
		new(Action),
	).Next(
		new(Verify),
	).Next(
		new(Invoke),
	)
	return handler
}
