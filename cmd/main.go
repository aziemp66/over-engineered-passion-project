package main

import (
	"fmt"
	pkg_http "simple-golang-rest-server/internal/pkg/http"
	"simple-golang-rest-server/router"

	"github.com/gin-gonic/gin"
)

const PORT int = 8000

func main() {
	srv := pkg_http.NewHTTPServer(gin.DebugMode)

	router.BindRouter(srv)

	srv.Run(fmt.Sprintf(":%d", PORT))
}
