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
)

type MvcErr int

const (
	// ServerErr 服务内部错误
	ServerErr MvcErr = iota + 599
	// APINotExist 接口不存在
	APINotExist
	// APIMappingCode 找不到code映射的api
	APIMappingCode
	// APIInjectFaild 接口注入失败
	APIInjectFaild
	// APIParemter api参数失败
	APIParemterFaild
)
