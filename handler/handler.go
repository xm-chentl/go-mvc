package handler

import "github.com/xm-chentl/go-mvc"

// MainHandler 主处理入口
type MainHandler struct {
	baseHandler
}

// Execute 执行
func (m MainHandler) Execute(ctx mvc.IContext) {
	// todo: 当前handler方法体

	if m.nextHandler != nil {
		m.nextHandler.Execute(ctx)
	}
}

func New(handlers ...mvc.IHandler) mvc.IHandler {
	mainHandler := &MainHandler{}
	var nextHandler mvc.IHandler
	nextHandler = mainHandler
	for _, h := range handlers {
		nextHandler.Next(h)
		nextHandler = h
	}
	return mainHandler
}
