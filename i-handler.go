package mvc

// IHandler 处理链接口
type IHandler interface {
	Next(IHandler) IHandler
	Exec(IContext)
}
