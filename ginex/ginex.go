package ginex

import (
	"fmt"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/context"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/verify/validator"

	"github.com/gin-gonic/gin"
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
	verifyInst := validator.New()
	ginInst.POST("/", func(ctx *gin.Context) {
		c := context.New()
		c.Set(enum.CTX, newRoute(ctx))
		c.Set(enum.Verify, verifyInst)
		g.handler.Execute(c)
	})

	fmt.Println("port: ", port)
	ginInst.Run(fmt.Sprintf("localhost:%d", port))
}

func New() mvc.IMvc {
	return &ginex{}
}
