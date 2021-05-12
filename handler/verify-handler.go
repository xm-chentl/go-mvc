package handler

import (
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/verify"
)

// VerifyHandler 校验处理
type VerifyHandler struct {
	baseHandler
}

func (h VerifyHandler) Execute(ctx mvc.IContext) {
	// todo: 方法体
	routeCtx := ctx.Get(enum.CTX).(mvc.IRoute)
	apiInstance := ctx.Get(enum.API).(mvc.IApi)
	verifyInstance := ctx.Get(enum.Verify).(verify.IVerify)
	routeCtx.Bind(apiInstance)
	if ok := verifyInstance.Execute(apiInstance); !ok {
		h.Error(ctx, enum.APIParemterFaild, "api parameter verify faild")
	}

	if h.nextHandler != nil {
		h.nextHandler.Execute(ctx)
	}
}
