package hello_handler_test

import (
	hello_controller "github.com/aziemp66/over-engineered-passion-project/internal/modules/hello/handler"
	pkg_http "github.com/aziemp66/over-engineered-passion-project/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

func setupTest() *gin.Engine {
	srv := pkg_http.NewHTTPServer(gin.TestMode)

	hello_controller.BindHelloHandler(srv.Group("/hello"))
	return srv
}
