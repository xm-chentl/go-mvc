package ginex

import (
	"fmt"
	"net/http"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/context"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/errorex"
	"github.com/xm-chentl/go-mvc/handler"
	"github.com/xm-chentl/go-mvc/verify/validator"

	"github.com/gin-gonic/gin"
)

type Option func(g *gin.Engine)

type ginex struct {
	options []Option
}

func (g ginex) Run(port int) {
	ginInstance := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	for _, optionFunc := range g.options {
		optionFunc(ginInstance)
	}

	if err := ginInstance.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}

func New(options ...Option) mvc.IMvc {
	return &ginex{
		options: options,
	}
}

func NewDefaultPost() Option {
	return func(ginInst *gin.Engine) {
		verifyInst := validator.New()
		handler := handler.Default()
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
			c.Set(enum.CTX, NewRoute(ctx))
			c.Set(enum.Verify, verifyInst)
			handler.Execute(c)
		})
	}
}
