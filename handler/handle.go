package handler

import (
	"github.com/xm-chentl/go-mvc"
	actionresult "github.com/xm-chentl/go-mvc/action-result"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/err"
	"github.com/xm-chentl/go-mvc/metadata"
)

// Handle 处理节点（初始化）
type Handle struct {
	nextHandler mvc.IHandler
}

// Next 下一节点
func (b *Handle) Next(handler mvc.IHandler) mvc.IHandler {
	b.nextHandler = handler
	return b.nextHandler
}

// Exec 执行
func (b Handle) Exec(ctx mvc.IContext) {
	if ok := ctx.Has(enum.Err); ok {
		errMsg := ctx.Get(enum.Err).(string)
		b.respError(ctx, errMsg)
		return
	}
	if !metadata.IsInitialize() {
		b.respError(ctx, err.APIInitErr)
		return
	}
	if b.nextHandler != nil {
		b.nextHandler.Exec(ctx)
	}
}

func (b Handle) respError(ctx mvc.IContext, errMsg string) {
	route := ctx.Get(enum.CTX).(mvc.IRoute)
	route.Response(
		actionresult.Error(500, errMsg),
	)
}
