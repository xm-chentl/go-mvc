package mvc

// IMvc 接口
type IMvc interface {
	SetHandle(IHandler) IMvc
	Run(port int)
}
