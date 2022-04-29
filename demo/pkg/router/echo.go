package router

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	_ "helloWorld/docs/mcy"
	"helloWorld/pkg/api"
	"helloWorld/pkg/config/config"
	"os"
	"os/signal"
	"time"
)

type Echo struct {
	echo         *echo.Echo
	serverConfig *config.ServerConfig
	log          *zap.Logger
}

func NewServer(api *api.Api, c *config.Config, log *zap.Logger) *Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: c.Server.Cors,
		AllowMethods: []string{fasthttp.MethodGet, fasthttp.MethodPut, fasthttp.MethodPost, fasthttp.MethodDelete},
	}))

	e.POST("/api/v1/check", api.ExampleApi.Check)
	log.Info("POST api/v1/check route defined")

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	log.Info("GET /metrics route defined")

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	log.Info("GET /swagger/* route defined")

	return &Echo{e, &c.Server, log}
}

func (e *Echo) RunEcho() {
	if err := e.echo.Start(e.serverConfig.Port); err != nil {
		e.echo.Logger.Fatal("Shutting down the server", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.echo.Shutdown(ctx); err != nil {
		e.echo.Logger.Fatal(err)
	}
}
