package handler

import (
	"reflect"
	"strings"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/ioc"
)

// InjectHandler 接口注入处理
type InjectHandler struct {
	baseHandler
}

// Exceute 处理
func (h InjectHandler) Execute(ctx mvc.IContext) {
	// todo: 方法体
	routeCtx := ctx.Get(enum.CTX).(mvc.IRoute)
	apiInstance := ctx.Get(enum.API).(mvc.IApi)
	err := ioc.Inject(&apiInstance, func(field reflect.StructField) interface{} {
		if strings.Contains(field.Type.Name(), "IRoute") {
			return routeCtx
		}

		return nil
	})
	if err != nil {
		h.Error(ctx, enum.APIInjectFaild, err.Error())
		return
	}

	if h.nextHandler != nil {
		h.nextHandler.Execute(ctx)
	}
}
