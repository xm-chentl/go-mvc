package mvc

import "net/http"

// IRoute 路由核心
type IRoute interface {
	Bind(interface{})
	Request() *http.Request
	Response() *http.Response
	Result(interface{})
}
