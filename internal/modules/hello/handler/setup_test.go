package hello_handler_test

import (
	hello_controller "simple-golang-rest-server/internal/modules/hello/handler"
	pkg_http "simple-golang-rest-server/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

func setupTest() *gin.Engine {
	srv := pkg_http.NewHTTPServer(gin.TestMode)

	hello_controller.BindHelloHandler(srv.Group("/hello"))
	return srv
}
