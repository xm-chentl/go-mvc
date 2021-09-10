package ginex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/xm-chentl/go-mvc"

	"github.com/gin-gonic/gin"
)

type ginRoute struct {
	ctx *gin.Context
}

func (g ginRoute) Bind(arg interface{}) {
	if g.ctx.Request.ContentLength > 0 {
		bodyByte, err := ioutil.ReadAll(g.ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(bodyByte, arg); err != nil {
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

func NewRoute(ctx *gin.Context) mvc.IRoute {
	return &ginRoute{
		ctx: ctx,
	}
}
