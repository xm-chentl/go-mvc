package enum

type ContextValue int

const (
	// CTX 路由组件上下文内容
	CTX ContextValue = iota
	// Code 接口编号
	Code
	// API 接口实现
	API
	// Result 响应
	Result
	// Verify 验证器
	Verify
	// Error is 错误
	Error
)
