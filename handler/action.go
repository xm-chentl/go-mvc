package handler

import (
	"fmt"

	"github.com/xm-chentl/go-mvc/metadata"
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/err"
)

// Action 行为路由
type Action struct {
	Handle
}

// Exec 执行
func (a Action) Exec(ctx mvc.IContext) {
	defer a.Handle.Exec(ctx)

	actionName := ctx.Get(enum.ActionName).(string)
	if actionName == "" {
		ctx.Set(enum.Err, err.APIPathErr)
		return
	}

	serviceName := ctx.Get(enum.ServiceName).(string)
	service := metadata.Get(serviceName)
	if _, ok := service.NameOfAction[actionName]; !ok {
		ctx.Set(
			enum.Err,
			fmt.Sprintf("%s action(%s)", err.APIServiceNotExist, actionName),
		)
		return
	}
}
