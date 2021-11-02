package common

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gitlab.secoder.net/bauhinia/qanda-schema/ent"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
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

func (c *Context) Verify(token string) (*jwt.RegisteredClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(_ *jwt.Token) (interface{}, error) { return c.Key, nil })
	if err != nil {
		return nil, err
	}
	return parsed.Claims.(*jwt.RegisteredClaims), nil
}

func (c *Context) Gensig(id string) (string, error) {
	return tencentyun.GenUserSig(1400586942, "b1c5ac2dd23bc7556ab94e23d2735806641f8d7fb3be28b779b66fe1672e6dd6", id, 86400*180)
}
