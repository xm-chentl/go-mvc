package handler

import (
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/errorex"
)

// CodeHandler 接口链处理
type CodeHandler struct {
	baseHandler
}

// Execute 执行
func (h CodeHandler) Execute(ctx mvc.IContext) {
	// todo: 方法体
	routeCtx := ctx.Get(enum.CTX).(mvc.IRoute)
	code := routeCtx.Request().Header.Get("code")
	if code == "" {
		// todo: 异常
		h.Error(ctx, errorex.APINotExist, "request header code is empty")
		return
	}
	ctx.Set(enum.Code, code)

	if h.nextHandler != nil {
		h.nextHandler.Execute(ctx)
	}
}
