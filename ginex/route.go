package ginex

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xm-chentl/go-mvc"
)

type ginRoute struct {
	ctx *gin.Context
}

func (g ginRoute) Bind(req interface{}) {
	if g.ctx.Request.ContentLength > 0 {
		if err := g.ctx.Bind(req); err != nil {
			// todo: 暂时是抛异常
			panic(err)
		}
	}
}

func (g *ginRoute) Request() *http.Request {
	return g.ctx.Request
}

func (g *ginRoute) Response() *http.Response {
	return nil
}

func (g *ginRoute) Result(data interface{}) {
	g.ctx.JSON(http.StatusOK, data)
}

func newRoute(ctx *gin.Context) mvc.IRoute {
	return &ginRoute{
		ctx: ctx,
	}
}
