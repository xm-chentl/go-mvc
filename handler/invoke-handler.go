package handler

import (
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
)

// InjectHandler 注入处理
type InvokeHandler struct {
	baseHandler
}

// Execute 处理
func (h InvokeHandler) Execute(ctx mvc.IContext) {
	// desc: 方法体
	apiInstance := ctx.Get(enum.API).(mvc.IApi)
	data, err := apiInstance.Call()
	ctx.Set(enum.Result, data)
	ctx.Set(enum.Error, err)
	// todo: 释放ctx

	if h.nextHandler != nil {
		h.nextHandler.Execute(ctx)
	}
}
