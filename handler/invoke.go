package handler

import (
	"reflect"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/err"
	"github.com/xm-chentl/go-mvc/metadata"
)

// Invoke 执行
type Invoke struct {
	Handle
}

// Exec 执行
func (i Invoke) Exec(ctx mvc.IContext) {
	defer i.Handle.Exec(ctx)

	serviceName := ctx.Get(enum.ServiceName).(string)
	actionName := ctx.Get(enum.ActionName).(string)
	service := metadata.Get(serviceName)
	args := ctx.Get(enum.Parameters).([]reflect.Value)
	actionResult := service.NameOfAction[actionName].Method.Func.Call(args)[0].Interface()
	resp := actionResult.(mvc.IActionResult)
	if resp == nil {
		ctx.Set(enum.Err, err.APIReturnErr)
		return
	}

	route := ctx.Get(enum.CTX).(mvc.IRoute)
	route.Response(resp.Exec())
}
