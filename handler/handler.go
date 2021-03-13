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

// New 设置责任链节点，并生成一个主节点
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

// Default 默认格式责任链
func Default() mvc.IHandler {
	return New(
		new(handler.CodeHandler),
		new(handler.APIHandler),
		new(handler.InjectHandler),
		new(handler.InvokeHandler),
		new(handler.ResultHandler),
	)
}
