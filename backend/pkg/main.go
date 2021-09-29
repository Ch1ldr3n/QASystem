package qanda

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	_ "gitlab.secoder.net/bauhinia/qanda/backend/pkg/docs"
	"gitlab.secoder.net/bauhinia/qanda/backend/pkg/question"
)

var Echo *echo.Echo

// @title Q&A API
// @version 1.0

// @host qanda-bauhinia.app.secoder.net
// @BasePath /
func init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/docs/*", echoSwagger.WrapHandler)
	v1 := e.Group("/v1")
	question.Register(v1)
	Echo = e
}
