package handler

import (
	"fmt"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/actionresult"
	"github.com/xm-chentl/go-mvc/enum"
)

type baseHandler struct {
	nextHandler mvc.IHandler
}

func (h *baseHandler) Next(handler mvc.IHandler) mvc.IHandler {
	h.nextHandler = handler
	return h.nextHandler
}

func (h *baseHandler) Error(ctx mvc.IContext, code enum.MvcErr, msg string) {
	// desc: 断开链
	h.nextHandler = nil
	routeCtx := ctx.Get(enum.CTX).(mvc.IRoute)
	routeCtx.Result(
		actionresult.Alert(
			int(code),
			msg,
		).Execute(),
	)
}

func (h *baseHandler) Errorf(ctx mvc.IContext, code enum.MvcErr, format string, args ...interface{}) {
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	h.Error(ctx, code, msg)
}
