package enum

// Value 枚举值
type Value int

const (
	// CTX 上下文内容
	CTX Value = iota
	// ServerName 服务名
	ServerName
	// ServiceName 服务名
	ServiceName
	// ActionName 行为名
	ActionName
	// Parameters 参数
	Parameters
	// RespFunc 响应回调
	RespFunc
	// Err 错误
	Err
	// TraceID 追溯标识
	TraceID
)
