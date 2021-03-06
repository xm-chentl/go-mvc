package handler

import (
	"reflect"
	"strings"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/metadata""
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
	// desc: 属性注入
	routeCtx.Bind(apiInstance)
	// desc: 组件注入
	rt := reflect.TypeOf(apiInstance).Elem()
	rv := reflect.ValueOf(apiInstance).Elem()
	containerErr := false
	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		// desc: 用于区分属性或者组件注入
		if field.Type.Kind() == reflect.Interface {
			// desc: 注入
			if strings.Contains(field.Type.Name(), "IRoute") {
				// desc: 注入上下文内容
				rv.Field(i).Set(
					reflect.ValueOf(routeCtx),
				)
				continue
			}

			key := strings.ToLower(field.Name)
			if container.Has(key) {
				// desc: 注入组件
				rv.Field(i).Set(
					reflect.ValueOf(
						container.Get(key),
					),
				)
				continue
			}
			containerErr = true
		}
	}
	if containerErr {
		// desc: 注入失败
		h.Error(ctx, enum.APIInjectFaild, "container inject faild")
		return
	}

	if h.nextHandler != nil {
		h.nextHandler.Execute(ctx)
	}
}
