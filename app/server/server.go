package server

import (
	"captcha-service/app/config"
	customMw "captcha-service/app/middleware/logging"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func InitEcho() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.CORS(),
		// middleware.Logger(),
		customMw.Logging(),
	)

	e.GET("/", func(c echo.Context) error {
		return ResponseOK(c, "captcha service API is running...", nil)
	})

	return e
}

// Start server
func Start(e *echo.Echo) {
	var (
		addr              = fmt.Sprintf(":%v", config.GetConfig().HTTPPort)
		HTTPServerTimeout = config.GetConfig().HTTPServerTimeOut
		readTime          = 1 * HTTPServerTimeout
		writeTime         = 20 * HTTPServerTimeout
	)

	s := &http.Server{
		Addr:         addr,
		ReadTimeout:  time.Second * time.Duration(readTime),
		WriteTimeout: time.Second * time.Duration(writeTime),
	}

	// Start server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("Shutting down the server ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
