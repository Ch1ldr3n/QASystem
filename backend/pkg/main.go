package qanda

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/swaggo/echo-swagger"
	"gitlab.secoder.net/bauhinia/qanda-schema/ent"
	"gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
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
func New(serve string, storage string, database string) *echo.Echo {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	db, err := ent.Open(storage, database)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Schema.Create(context.Background()); err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &common.Context{Context: c, DBField: db}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", serve)
	e.GET("/docs/*", echoSwagger.WrapHandler)
	v1 := e.Group("/v1")
	user.Register(v1.Group("/user"))
	question.Register(v1.Group("/question"))
	return e
}
