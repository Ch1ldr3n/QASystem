package qanda

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUser(t *testing.T) {
	e := New("/var/empty", "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", "super-secret-key")

	req := httptest.NewRequest(http.MethodPost, "/v1/user/register", bytes.NewBufferString(`
{
	"username": "testuser",
	"password": "testpassword"
}
    `))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("register failed")
	}

	req = httptest.NewRequest(http.MethodPost, "/v1/user/login", bytes.NewBufferString(`
{
	"username": "testuser",
	"password": "testpassword"
}
    `))
	req.Header.Add("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("login failed")
	}
	resp := new(struct {
		Token string `json:"token"`
	})
	err := json.NewDecoder(rec.Body).Decode(resp)
	if err != nil {
		t.Fatal(err)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/user/info", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", resp.Token)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("get user info failed")
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/user/filter?username=test", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("filter user failed")
	}
}

func TestQuestionSubmit(t *testing.T) {
	e := New("/var/empty", "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", "super-secret-key")

	{	//create user1 -- supposing his id is 1
		req := httptest.NewRequest(http.MethodPost, "/v1/user/register", bytes.NewBufferString(`
{
	"username": "testuse1",
	"password": "testpassword"
}
		`))
		req.Header.Add("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
	}
	{	//create user2	-- supposing his id is 2
		req := httptest.NewRequest(http.MethodPost, "/v1/user/register", bytes.NewBufferString(`
{
	"username": "testuser2",
	"password": "testpassword"
}
		`))
		req.Header.Add("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/question/submit", bytes.NewBufferString(`
{
	"price": 100,
	"title": "test title",
	"content":"test content",
	"questionerid":1,
	"answererid":2
}
	`))
	// Note: if generation pattern of user.ID changes, so shall the line above do as well.
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Result().StatusCode != http.StatusOK {
		t.Fatal("submit question failed")
	}
}