package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
	"gitlab.secoder.net/bauhinia/qanda-schema/ent"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
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

func (c *Context) RequestTIM(svc string, cmd string, body interface{}) error {
	query := make(url.Values)
	query.Set("sdkappid", "1400586942")
	query.Set("identifier", "administrator")
	sig, err := c.Gensig("administrator")
	if err != nil {
		return err
	}
	query.Set("usersig", sig)
	query.Set("contenttype", "json")
	query.Set("random", strconv.Itoa(int(rand.Uint32())))
	u := url.URL{
		Scheme:   "https",
		Host:     "console.tim.qq.com",
		Path:     fmt.Sprintf("/v4/%s/%s", svc, cmd),
		RawQuery: query.Encode(),
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err = enc.Encode(body)
	if err != nil {
		return err
	}
	resp, err := http.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed")
	}
	return nil
}
