package ginex

import (
	"fmt"
	"net/http"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/context"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/errorex"
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
	ginInst.POST("/:server/:action", func(ctx *gin.Context) {
		defer func() {
			if recoverErr := recover(); recoverErr != nil {
				ctx.JSON(
					http.StatusOK,
					map[string]interface{}{
						"err":  errorex.ServerErr,
						"data": map[string]interface{}{},
					},
				)
			}
		}()

		c := context.New()
		c.Set(enum.Code, fmt.Sprintf("/%s/%s", ctx.Param("server"), ctx.Param("action")))
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
