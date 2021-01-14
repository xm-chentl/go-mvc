package mvc

import "net/http"

// IRoute 路由接口
type IRoute interface {
	Bind(interface{})
	Request() *http.Request
	Response(interface{})
}
