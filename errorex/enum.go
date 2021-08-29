package errorex

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
