package qanda

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo/v4"
	"strconv"
	"fmt"
)

// auxiliary functions for auxiliary functions

func getTokenFromBody(rec *httptest.ResponseRecorder, t *testing.T) string {
	resp := new(struct {
		Token string `json:"token"`
	})
	err := json.NewDecoder(rec.Body).Decode(resp)
	if t != nil && err != nil {
		t.Fatal(err)
	}
	return resp.Token
} 

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

	if t != nil && rec.Result().StatusCode != http.StatusOK {
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

	if t != nil && rec.Result().StatusCode != http.StatusOK {
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

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("get user info failed")
	}

	return rec
}

func auxUserFilter(e *echo.Echo, t *testing.T, query string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/user/filter"+query, nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("filter user failed")
	}

	return rec
}

func auxQuestionSubmit(e *echo.Echo, t *testing.T, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/submit", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("submit question failed")
	}

	return rec
}

func auxQuestionPay(e *echo.Echo, t *testing.T, questionid int, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/pay", bytes.NewBufferString(`
{
	"questionid": `+strconv.Itoa(questionid)+`
}
    `))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question pay failed")
	}

	return rec
}

func auxQuestionQuery(e *echo.Echo, t *testing.T, questionid int) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/question/"+strconv.Itoa(questionid), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question query failed")
	}

	return rec
}

func auxQuestionList(e *echo.Echo, t *testing.T) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/question/list", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question list failed")
	}

	return rec
}

func auxQuestionMine(e *echo.Echo, t *testing.T, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/question/mine", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question mine failed")
	}

	return rec
}

func auxQuestionAccept(e *echo.Echo, t *testing.T, questionid int, choice bool, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/accept", bytes.NewBufferString(`
{
	"questionid": `+strconv.Itoa(questionid)+`,
	"choice": `+fmt.Sprintf(`%t`,choice)+`
}
    `))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question accept failed")
	}

	return rec
}

func auxQuestionClose(e *echo.Echo, t *testing.T, questionid int, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/close", bytes.NewBufferString(`
{
	"questionid": `+strconv.Itoa(questionid)+`
}
    `))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question close failed")
	}

	return rec
}

// test functions

func TestUser(t *testing.T) {
	e := New("/var/empty", "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", "super-secret-key")

	auxUserRegister(e, t, "testuser", "testpassword")
	rec := auxUserLogin(e, t, "testuser", "testpassword")
	token := getTokenFromBody(rec, t)
	auxUserInfo(e, t, token)
	auxUserFilter(e, t, "?username=testuser")
}

func TestQuestion(t *testing.T) {
	e := New("/var/empty", "sqlite3", "file:ent2?mode=memory&cache=shared&_fk=1", "super-secret-key")

	rec := auxUserRegister(e, t, "user1", "pass")
	token1 := getTokenFromBody(rec, t)
	rec = auxUserRegister(e, t, "user2", "pass")
	token2 := getTokenFromBody(rec, t)

	auxQuestionSubmit(e, t, `
{
	"price": 0,
	"title": "test title",
	"content":"test content",
	"questionerid":1,
	"answererid":2
}
	`)

	auxQuestionPay(e, t, 1, token1)
	auxQuestionQuery(e, t, 1)
	auxQuestionList(e, t)
	auxQuestionMine(e, t, token1)

	auxQuestionAccept(e, t, 1, true, token2)
	auxQuestionClose(e, t, 1, token1)
}