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
)

type MvcErr int

const (
	// APINotExist 接口不存在
	APINotExist MvcErr = iota + 601
	// APIMappingCode 找不到code映射的api
	APIMappingCode
	// APIInjectFaild 接口注入失败
	APIInjectFaild
	// APIParemter api参数失败
	APIParemterFaild
)
