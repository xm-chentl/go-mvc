package handler

import (
	"reflect"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/metadata"
)

// Verify 验证
type Verify struct {
	Handle
}

// Exec 执行
func (v Verify) Exec(ctx mvc.IContext) {
	defer v.Handle.Exec(ctx)

	serverName := ctx.Get(enum.ServiceName).(string)
	actionName := ctx.Get(enum.ActionName).(string)
	server := metadata.Get(serverName)
	methodOfAction := server.NameOfAction[actionName]
	route := ctx.Get(enum.CTX).(mvc.IRoute)
	rvs := []reflect.Value{reflect.ValueOf(server.Instance)}
	for i := 0; i < len(methodOfAction.Parameters); i++ {
		rt := methodOfAction.Method.Type.In(i)
		switch rt.Kind() {
		case reflect.Interface:
			rvs = append(rvs, reflect.ValueOf(route))
			break
		case reflect.Ptr:
			reqOfArge := reflect.New(rt.Elem()).Interface()
			route.Bind(reqOfArge)
			rvs = append(rvs, reflect.ValueOf(reqOfArge))
			break
		}
	}
	ctx.Set(enum.Parameters, rvs)
}
