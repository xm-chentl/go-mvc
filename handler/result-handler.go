package handler

import (
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/errorex"
)

// ResultHandler 响应处理
type ResultHandler struct {
	baseHandler
}

// Exceute 执行
func (h ResultHandler) Execute(ctx mvc.IContext) {
	// desc: 响应
	defer func() {
		if h.nextHandler != nil {
			h.nextHandler.Execute(ctx)
		}
	}()

	apiResult := make(map[string]interface{})
	err := ctx.Get(enum.Error)
	if err != nil {
		if e, ok := err.(errorex.MvcError); ok {
			apiResult["err"] = e.Code()
			apiResult["data"] = e.Error()
		} else {
			apiResult["err"] = errorex.ServerErr
			apiResult["data"] = map[string]interface{}{}
		}
	} else {
		apiResult["err"] = 0
		apiResult["data"] = ctx.Get(enum.Result)
	}
	routeCtx := ctx.Get(enum.CTX).(mvc.IRoute)
	routeCtx.Result(apiResult)
}
