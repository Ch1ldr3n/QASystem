package qanda

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	_ "gitlab.secoder.net/bauhinia/qanda/backend/pkg/docs"
	"gitlab.secoder.net/bauhinia/qanda/backend/pkg/question"
	"gitlab.secoder.net/bauhinia/qanda/backend/pkg/user"
	"net/http"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// @title Q&A API
// @version 1.0

// @host qanda-bauhinia.app.secoder.net
// @BasePath /
func New(serve string) *echo.Echo {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", serve)
	e.GET("/docs/*", echoSwagger.WrapHandler)
	v1 := e.Group("/v1")
	user.Register(v1.Group("/user"))
	question.Register(v1.Group("/question"))
	return e
}
