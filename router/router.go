package router

import (
	hello_handler "simple-golang-rest-server/internal/modules/hello/handler"

	"github.com/gin-gonic/gin"
)

func BindRouter(router *gin.Engine) {

	hello_handler.BindHelloHandler(router.Group("/hello"))
}
