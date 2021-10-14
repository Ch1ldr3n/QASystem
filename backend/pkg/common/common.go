package common

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gitlab.secoder.net/bauhinia/qanda-schema/ent"
)

type Context struct {
	echo.Context
	DBField *ent.Client
	Key     []byte
}

func (c *Context) DB() *ent.Client {
	return c.DBField
}

func (c *Context) Sign(username string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Subject: username,
	}).SignedString(c.Key)
}

func (c *Context) Verify() *ent.Client {
	return c.DBField
}
