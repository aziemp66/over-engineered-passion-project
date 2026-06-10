package router

import (
	hello_handler "github.com/aziemp66/over-engineered-passion-project/internal/modules/hello/handler"

	"github.com/gin-gonic/gin"
)

func BindRouter(router *gin.Engine) {

	hello_handler.BindHelloHandler(router.Group("/hello"))
}
