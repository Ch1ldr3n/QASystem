package qanda

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo/v4"
)

// auxiliary functions
func auxUserRegister(e *echo.Echo, t *testing.T, name string, password string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/user/register", bytes.NewBufferString(`
{
	"username": "`+name+`",
	"password": "`+password+`"
}
    `))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("register failed")
	}

	return rec
}

func auxUserLogin(e *echo.Echo, t *testing.T, name string, password string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/user/login", bytes.NewBufferString(`
{
	"username": "`+name+`",
	"password": "`+password+`"
}
    `))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("login failed")
	}

	return rec
}

func auxUserInfo(e *echo.Echo, t *testing.T, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/user/info", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("get user info failed")
	}

	return rec
}

func auxUserFilter(e *echo.Echo, t *testing.T, query string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/user/filter"+query, nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("filter user failed")
	}

	return rec
}

func auxQuestionSubmit(e *echo.Echo, t *testing.T, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/submit", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("submit question failed")
	}

	return rec
}

// test functions

func TestUser(t *testing.T) {
	e := New("/var/empty", "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", "super-secret-key")

	auxUserRegister(e, t, "testuser", "testpassword")
	rec := auxUserLogin(e, t, "testuser", "testpassword")
	resp := new(struct {
		Token string `json:"token"`
	})
	err := json.NewDecoder(rec.Body).Decode(resp)
	if err != nil {
		t.Fatal(err)
	}
	auxUserInfo(e, t, resp.Token)
	auxUserFilter(e, t, "?username=testuser")
}

func TestQuestion(t *testing.T) {
	e := New("/var/empty", "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", "super-secret-key")

	auxUserRegister(e, t, "user1", "pass")
	auxUserRegister(e, t, "user2", "pass")

	auxQuestionSubmit(e, t, `
{
	"price": 100,
	"title": "test title",
	"content":"test content",
	"questionerid":1,
	"answererid":2
}
	`)
}