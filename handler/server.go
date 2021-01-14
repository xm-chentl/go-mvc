package handler

import (
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/err"
)

// Server 服务服务
type Server struct {
	Handle
}

// Exec 执行
func (s Server) Exec(ctx mvc.IContext) {
	defer s.Handle.Exec(ctx)

	serverName := ctx.Get(enum.ServerName)
	if serverName == "" {
		ctx.Set(enum.Err, err.APIServerNotExist)
		return
	}
}
