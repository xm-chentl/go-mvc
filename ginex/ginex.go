package ginex

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/context"
	"github.com/xm-chentl/go-mvc/enum"
)

type ginex struct {
	handler mvc.IHandler
}

func (g *ginex) AddHandler(handler mvc.IHandler) mvc.IMvc {
	g.handler = handler
	return g
}

func (g ginex) Run(port int) {
	ginInst := gin.Default()
	ginInst.POST("/", func(ctx *gin.Context) {
		c := context.New()
		c.Set(enum.CTX, newRoute(ctx))
		// todo: 不可开协程
		g.handler.Execute(c)
	})

	fmt.Println("port: %d", port)
	ginInst.Run(fmt.Sprintf("localhost:%d", port))
}

func New() mvc.IMvc {
	return &ginex{}
}