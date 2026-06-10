package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	pkg_http "github.com/aziemp66/over-engineered-passion-project/internal/pkg/http"
	pkg_logger "github.com/aziemp66/over-engineered-passion-project/internal/pkg/logger"
	"github.com/aziemp66/over-engineered-passion-project/router"

	"github.com/gin-gonic/gin"
)

const PORT int = 8000

func main() {
	srv := pkg_http.NewHTTPServer(gin.DebugMode)
	pkg_logger.InitLogger(gin.Mode(), "log/app.log")

	router.BindRouter(srv)

	// Start server
	go func() {
		pkg_logger.Info(fmt.Sprintf("Server Listening on : %d", PORT))
		if err := srv.Run(fmt.Sprintf(":%d", PORT)); err != nil {
			pkg_logger.Fatal("Fail to start http server")
		}
	}()

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done

	pkg_logger.Info(fmt.Sprintf("Server gracefully shutdown on %s", time.Now().Format(time.RFC3339)))
}
