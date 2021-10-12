package common

import (
	"github.com/labstack/echo/v4"
	"gitlab.secoder.net/bauhinia/qanda-schema/ent"
)

type Context struct {
	echo.Context
	DBField *ent.Client
}

func (c *Context) DB() *ent.Client {
	return c.DBField
}
