package qanda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

//
// Auxiliary functions for Auxiliary functions
//

func GetEchoTestEnv(filename string) *echo.Echo {
	return New("/var/empty", "sqlite3", "file:"+filename+"?mode=memory&cache=shared&_fk=1", "super-secret-key")
}

func GetIdTokenFromRec(rec *httptest.ResponseRecorder, t *testing.T) (string, int) {
	resp := new(struct {
		Token string `json:"token"`
		ID    int    `json:"id"`
	})
	err := json.NewDecoder(rec.Body).Decode(resp)
	if t != nil && err != nil {
		t.Fatal(err)
	}
	return resp.Token, resp.ID
}

func GetQuestionIdFromSubmit(rec *httptest.ResponseRecorder, t *testing.T) int {
	resp := new(struct {
		QuestionID int `json:"questionid"`
	})
	err := json.NewDecoder(rec.Body).Decode(resp)
	if t != nil && err != nil {
		t.Fatal(err)
	}
	return resp.QuestionID
}

//
// Auxiliary functions
//

func AuxUserRegister(e *echo.Echo, t *testing.T, name string, password string) *httptest.ResponseRecorder {
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

func AuxUserLogin(e *echo.Echo, t *testing.T, name string, password string) *httptest.ResponseRecorder {
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

func AuxUserGensig(e *echo.Echo, t *testing.T, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/user/gensig", nil)
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("user signature generation failed")
	}

	return rec
}

func AuxUserInfo(e *echo.Echo, t *testing.T, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/user/info", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("Get user info failed")
	}

	return rec
}

func AuxUserEdit(e *echo.Echo, t *testing.T, token string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/user/edit", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("edit user failed")
	}

	return rec
}

func AuxUserFilter(e *echo.Echo, t *testing.T, token string, query string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/user/filter"+query, nil)
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("filter user failed")
	}

	return rec
}

func AuxQuestionSubmit(e *echo.Echo, t *testing.T, token string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/submit", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("submit question failed")
	}

	return rec
}

func AuxQuestionPay(e *echo.Echo, t *testing.T, questionid int, token string) *httptest.ResponseRecorder {
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

func AuxQuestionQuery(e *echo.Echo, t *testing.T, questionid int) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/question/"+strconv.Itoa(questionid), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question query failed")
	}

	return rec
}

func AuxQuestionList(e *echo.Echo, t *testing.T) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/question/list", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question list failed")
	}

	return rec
}

func AuxQuestionMine(e *echo.Echo, t *testing.T, token string) *httptest.ResponseRecorder {
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

func AuxQuestionAccept(e *echo.Echo, t *testing.T, questionid int, choice bool, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/accept", bytes.NewBufferString(`
{
	"questionid": `+strconv.Itoa(questionid)+`,
	"choice": `+fmt.Sprintf(`%t`, choice)+`
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

func AuxQuestionClose(e *echo.Echo, t *testing.T, questionid int, token string) *httptest.ResponseRecorder {
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

func AuxQuestionCancel(e *echo.Echo, t *testing.T, questionid int, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/cancel", bytes.NewBufferString(`
{
	"questionid": `+strconv.Itoa(questionid)+`
}
    `))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question cancel failed")
	}

	return rec
}

//
// test functions
//

// User:

func TestUser(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	AuxUserRegister(e, t, "user1", "testpassword")
	rec := AuxUserLogin(e, t, "user1", "testpassword")
	token1, userid1 := GetIdTokenFromRec(rec, t)
	rec = AuxUserRegister(e, t, "user2", "testpassword")
	token2, _ := GetIdTokenFromRec(rec, t)
	
	AuxUserInfo(e, t, token1)
	AuxUserGensig(e, t, token1)
	AuxUserEdit(e, t, token1, `
{
	"email":"hello",
	"phone":"12345678",
	"answerer":true,
	"price":-100,
	"profession":"Geschichte"
}
	`)
	AuxUserFilter(e, t, token2, "?id="+strconv.Itoa(userid1)+"&username=user1&email=hello&phone=12345678&answerer=true&priceUpperBound=1000&priceLowerBound=-1000&profession=Geschichte")
}

// Register: bad json
func TestUserRegisterX1(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	if rec := AuxUserRegister(e, nil, "userX\"!!", "testpassword"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user register allows bad json")
	}
}

// Register: no password
func TestUserRegisterX2(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	if rec := AuxUserRegister(e, nil, "userX\"}", ""); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user register allows no password")
	}
}

// Register: repeated register
func TestUserRegisterX3(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	if rec := AuxUserRegister(e, nil, "user1", "testpassword"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user register allows repeated register")
	}
}

// Login: bad json
func TestUserLoginX1(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	if rec := AuxUserLogin(e, nil, "user1\"!!", "testpassword"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user login allows bad json")
	}
}

// Login: no password
func TestUserLoginX2(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	if rec := AuxUserLogin(e, nil, "user1\"}", "testpassword"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user login allows no password")
	}
}

// Login: incorrect password
func TestUserLoginX3(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	if rec := AuxUserLogin(e, nil, "user1", "hello"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user login allows incorrect password")
	}
}

// Info: token verification
func TestUserInfoX1(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	token, _ := GetIdTokenFromRec(AuxUserLogin(e, t, "user1", "testpassword"), t)
	if rec := AuxUserInfo(e, nil, token+"qwerty"); rec.Result().StatusCode != http.StatusForbidden {
		t.Fatal("user info allows incorrect token")
	}
}

// Info: inexistent user
func TestUserInfoX2(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	e1 := GetEchoTestEnv("entTestUserInfoX2")
	token, _ := GetIdTokenFromRec(AuxUserRegister(e1, t, "user1X", "testpassword"), t)
	if rec := AuxUserInfo(e, nil, token); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user info allows inexistent user")
	}
}

// Edit: bad json
func TestUserEditX1(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	token, _ := GetIdTokenFromRec(AuxUserLogin(e, t, "user1", "testpassword"), t)
	if rec := AuxUserEdit(e, nil, token, "{\"}"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user edit allows bad json")
	}
}

// Filter: bad query
func TestUserFilterX1(t *testing.T) {
	e := GetEchoTestEnv("entUser")
	if rec := AuxUserFilter(e, nil, "", "?answerer=trualse"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user filter allows bad query")
	}
}

// Question:

func TestQuestion(t *testing.T) {
	e := GetEchoTestEnv("entQuestion")

	// Create 2 users
	rec := AuxUserRegister(e, t, "user1", "pass")
	token1, userid1 := GetIdTokenFromRec(rec, t)
	rec = AuxUserRegister(e, t, "user2", "pass")
	token2, userid2 := GetIdTokenFromRec(rec, t)

	AuxUserEdit(e, t, token1, `
{
	"answerer":true,
	"price":100
}
	`)
	AuxUserEdit(e, t, token2, `
{
	"answerer":true,
	"price":10
}
	`)

	// Create 3 questions
	rec = AuxQuestionSubmit(e, t, token1, `
{
	"title": "test title1",
	"content":"test content1",
	"answererid":`+strconv.Itoa(userid2)+`
}
	`)
	questionid1 := GetQuestionIdFromSubmit(rec, t)
	rec = AuxQuestionSubmit(e, t, token1, `
{
	"title": "test title2",
	"content":"test content2",
	"answererid":`+strconv.Itoa(userid2)+`
}
	`)
	questionid2 := GetQuestionIdFromSubmit(rec, t)
	rec = AuxQuestionSubmit(e, t, token1, `
{
	"title": "test title3",
	"content":"test content3",
	"answererid":`+strconv.Itoa(userid2)+`
}
	`)
	questionid3 := GetQuestionIdFromSubmit(rec, t)
	rec = AuxQuestionSubmit(e, t, token2, `
{
	"title": "test title4",
	"content":"test content4",
	"answererid":`+strconv.Itoa(userid1)+`
}
	`)
	// questionid4 := GetQuestionIdFromSubmit(rec, t)

	// Launch some global queries
	AuxQuestionQuery(e, t, questionid1)
	AuxQuestionList(e, t)
	AuxQuestionMine(e, t, token1)

	// Accept question no.1
	AuxQuestionPay(e, t, questionid1, token1)
	AuxQuestionAccept(e, t, questionid1, true, token2)

	// Reject question no.2
	AuxQuestionPay(e, t, questionid2, token1)
	AuxQuestionAccept(e, t, questionid2, false, token2)

	// Cancel question no.3
	AuxQuestionPay(e, t, questionid3, token1)
	AuxQuestionCancel(e, t, questionid3, token1)

	// Close question no.1
	AuxQuestionClose(e, t, questionid1, token1)
}

func TestQuestionX1(t *testing.T) {
	e := GetEchoTestEnv("entQuestion")
	token3, userid3 := GetIdTokenFromRec(AuxUserRegister(e, t, "user3", "pass"), t)
	AuxUserEdit(e, t, token3, `
{
	"answerer":true
}
	`)

	// Submit: questioning oneself
	if rec := AuxQuestionSubmit(e, nil, token3, `
{
	"title": "test titleX",
	"content":"test contentX",
	"answererid":`+strconv.Itoa(userid3)+`
}
	`); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question submit allows questioning oneself")
	}

	// Pay: repeated payment
	_, userid2 := GetIdTokenFromRec(AuxUserLogin(e, t, "user2", "pass"), t)
	questionid4 := GetQuestionIdFromSubmit(AuxQuestionSubmit(e, t, token3, `
{
	"title": "test title4",
	"content":"test content4",
	"answererid":`+strconv.Itoa(userid2)+`
}
	`), t)
	AuxQuestionPay(e, t, questionid4, token3)
	if rec := AuxQuestionPay(e, nil, questionid4, token3); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question pay allows repeated payment")
	}

	// Pay: cannot afford question	(abandoned)
	token1, userid1 := GetIdTokenFromRec(AuxUserLogin(e, t, "user1", "pass"), t)
	questionid5 := GetQuestionIdFromSubmit(AuxQuestionSubmit(e, t, token3, `
{
	"title": "test title5",
	"content":"test content5",
	"answererid":`+strconv.Itoa(userid1)+`
}
	`), t)
	// if rec := AuxQuestionPay(e, nil, questionid5, token3); rec.Result().StatusCode != http.StatusBadRequest {
	// 	t.Fatal("question pay allows illegal payment")
	// }

	// Pay: paying another person's question
	token4, _ := GetIdTokenFromRec(AuxUserRegister(e, t, "user4", "pass"), t)
	if rec := AuxQuestionPay(e, nil, questionid5, token4); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question pay allows paying others' questions")
	}

	// Accept: status not 'paid'
	if rec := AuxQuestionAccept(e, nil, questionid5, true, token1); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question accept allows wrong status")
	}

	// Close: status not 'accepted'
	if rec := AuxQuestionClose(e, nil, questionid5, token1); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question close allows wrong status")
	}

	// Cancel: double canceling
	AuxQuestionCancel(e, nil, questionid5, token1)
	if rec := AuxQuestionCancel(e, nil, questionid5, token1); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question cancel allows wrong status")
	}
}

func TestQuestionX2(t *testing.T) {
	e := GetEchoTestEnv("entQuestion")
	token1, _ := GetIdTokenFromRec(AuxUserLogin(e, t, "user1", "pass"), t)
	token2, userid2 := GetIdTokenFromRec(AuxUserLogin(e, t, "user2", "pass"), t)
	token3, _ := GetIdTokenFromRec(AuxUserLogin(e, t, "user3", "pass"), t)
	rec := AuxQuestionSubmit(e, t, token1, `
{
	"title": "test title6",
	"content":"test content6",
	"answererid":`+strconv.Itoa(userid2)+`
}
	`)
	questionid6 := GetQuestionIdFromSubmit(rec, t)
	AuxQuestionPay(e, t, questionid6, token1)
	
	// Accept: foreign interference
	if rec := AuxQuestionAccept(e, nil, questionid6, true, token3); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question accept allows foreign interference")
	}

	// Close: foreign interference
	AuxQuestionAccept(e, t, questionid6, true, token2)
	if rec := AuxQuestionAccept(e, nil, questionid6, true, token3); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question close allows foreign interference")
	}

	// Cancel: foreign interference
	if rec := AuxQuestionCancel(e, nil, questionid6, token3); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question cancel allows foreign interference")
	}
}

// Query: inexistent question
func TestQuestionQueryX1(t *testing.T) {
	e := GetEchoTestEnv("entQuestion")
	if rec := AuxQuestionQuery(e, nil, -1); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question query allows inexistent question")
	}
}

// Query: ill-formed ID
func TestQuestionQueryX2(t *testing.T) {
	e := GetEchoTestEnv("entQuestion")
	req := httptest.NewRequest(http.MethodGet, "/v1/question/123abc", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("question query allows ill-formed ID")
	}
}

// Unified test for bad json and invalid token verification
// 
// - the first string parameter of 'af' is its json item, the second its token
// - it's better to make sure that validation of json and token is done at the beginning of the api function, 
// - which means the 'Bind - BindHeaders - Validate - Verify' procedure
//

func AuxTestVerificationX(name string, t *testing.T, af func (*echo.Echo, *testing.T, string, string) *httptest.ResponseRecorder) {
	e := GetEchoTestEnv("entVerificationX"+name)
	e1 := GetEchoTestEnv("entVerificationX1"+name)
	token1, _ := GetIdTokenFromRec(AuxUserRegister(e, t, "userX", "pass"), t)
	token2, _ := GetIdTokenFromRec(AuxUserRegister(e1, t, "userXX", "pass"), t)

	// Bad json
	if rec := af(e, nil, ",***", token1); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("api allows bad json")
	}

	// No token
	if rec := af(e, nil, "", ""); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("api allows no token")
	}

	// Not verifiable
	if rec := af(e, nil, "", token1 + "qwerty"); rec.Result().StatusCode != http.StatusForbidden {
		t.Fatal("api allows bad verification")
	}

	// inexistent user
	if rec := af(e, nil, "", token2); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("api allows inexistent user")
	}
}

// func TestUserEditXv(t *testing.T) {
// 	AuxTestVerificationX("UserEdit", t, func (e *echo.Echo, t *testing.T, jsonitem string, token string) *httptest.ResponseRecorder{
// 		return AuxUserEdit(e, t, token, "{"+jsonitem+"}")
// 	})
// }
func TestQuestionSubmitXv(t *testing.T) {
	AuxTestVerificationX("QuestionSubmit", t, func (e *echo.Echo, t *testing.T, jsonitem string, token string) *httptest.ResponseRecorder{
		return AuxQuestionSubmit(e, t, token, `
		{
			"title": "test title1",
			"content":"test content1",
			"answererid":-1
			`+jsonitem+`
		}
			`)
	})
}
func TestQuestionPayXv(t *testing.T) {
	AuxTestVerificationX("QuestionPay", t, func (e *echo.Echo, t *testing.T, jsonitem string, token string) *httptest.ResponseRecorder{
		return AuxQuestionPay(e, t, -1, token)
	})
}