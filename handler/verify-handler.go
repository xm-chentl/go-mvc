package handler

import (
	"reflect"

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
	// desc: 属性注入
	routeCtx.Bind(apiInstance)
	// todo: 优化反射
	rt := reflect.TypeOf(apiInstance).Elem()
	rv := reflect.ValueOf(apiInstance).Elem()
	// desc: 获取参数
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.Type.Kind() != reflect.Interface {
			ver := verify.New(field.Tag.Lookup)
			if ok := ver.Execute(rv.Field(i).Interface()); !ok {
				h.Error(ctx, enum.APIParemterFaild, "api parameter verify faild")
				return
			}
		}
	}

	if h.nextHandler != nil {
		h.nextHandler.Execute(ctx)
	}
}
