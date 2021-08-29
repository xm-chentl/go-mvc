package handler

import (
	"fmt"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/actionresult"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/errorex"
)

type baseHandler struct {
	nextHandler mvc.IHandler
	nexts       []mvc.IHandler
}

func (h *baseHandler) Next(handler mvc.IHandler) mvc.IHandler {
	h.nexts = append(h.nexts, handler)
	return h
}

func (h *baseHandler) Error(ctx mvc.IContext, code errorex.MvcErr, msg string) {
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

func (h *baseHandler) Errorf(ctx mvc.IContext, code errorex.MvcErr, format string, args ...interface{}) {
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	h.Error(ctx, code, msg)
}

func (h *baseHandler) Execute(ctx mvc.IContext) {
	if h.nextHandler != nil {
		h.nextHandler.Execute(ctx)
	}
}
