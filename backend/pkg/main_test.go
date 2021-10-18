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
}
