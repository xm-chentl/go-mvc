package handler

import (
	"fmt"

	"gitee.com/love-yd/loveyd-lib-go/mvc"
	"gitee.com/love-yd/loveyd-lib-go/mvc/enum"
	"gitee.com/love-yd/loveyd-lib-go/mvc/err"
	"gitee.com/love-yd/loveyd-lib-go/mvc/metadata"
)

// Service 资源路由
type Service struct {
	Handle
}

// Exec 执行
func (s Service) Exec(ctx mvc.IContext) {
	defer s.Handle.Exec(ctx)

	serviceName := ctx.Get(enum.ServiceName).(string)
	if serviceName == "" {
		ctx.Set(enum.Err, err.APIPathErr)
		return
	}
	if ok := metadata.Has(serviceName); !ok {
		ctx.Set(
			enum.Err,
			fmt.Sprintf(
				"%s service(%s)",
				err.APIServiceNotExist,
				serviceName,
			),
		)
		return
	}
}
