package handler

import (
	"reflect"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/errorex"
	"github.com/xm-chentl/go-mvc/metadata"
)

// APIHandler 接口注入处理
type APIHandler struct {
	baseHandler
}

// Execute 执行
func (h APIHandler) Execute(ctx mvc.IContext) {
	// desc: 方法体
	code := ctx.Get(enum.Code).(string)
	if ok := metadata.Has(code); !ok {
		h.Errorf(ctx, errorex.APIMappingCode, "code (%s) mapping api not exist", code)
		return
	}

	api := metadata.Get(code)
	// desc: copy一个新实例，防止共享无数据
	rt := reflect.TypeOf(api).Elem()
	ctx.Set(enum.API, reflect.New(rt).Interface())

	if h.nextHandler != nil {
		h.nextHandler.Execute(ctx)
	}
}
