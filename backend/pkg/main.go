package qanda

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	_ "gitlab.secoder.net/bauhinia/qanda/backend/pkg/docs"
	"net/http"
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
	e.GET("/", index)
	e.GET("/docs/*", echoSwagger.WrapHandler)
	Echo = e
}

// @Summary Index
// @Router / [get]
func index(c echo.Context) error {
	return c.String(http.StatusOK, "index")
}
