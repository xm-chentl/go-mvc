package mvc

// IHandler 链式处理节点
type IHandler interface {
	Next(IHandler) IHandler
	Execute(IContext)
}
