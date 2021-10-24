package question

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuestionSubmit(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/submit", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	submit(c)
}

func TestQuestionQuery(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	query(c)
}
