package mvc

type IMvc interface {
	AddHandler(handler IHandler) IMvc
	// Method todo: 扩展其它协议 PUT, DELETE
	Run(port int)
}
