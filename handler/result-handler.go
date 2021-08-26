package handler

import (
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
)

// ResultHandler 响应处理
type ResultHandler struct {
	baseHandler
}

// Exceute 执行
func (h ResultHandler) Execute(ctx mvc.IContext) {
	// desc: 响应
	routeCtx := ctx.Get(enum.CTX).(mvc.IRoute)
	dataResult := ctx.Get(enum.Result)
	routeCtx.Result(dataResult)

	if h.nextHandler != nil {
		h.nextHandler.Execute(ctx)
	}
}
