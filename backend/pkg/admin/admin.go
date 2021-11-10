package admin

import "github.com/labstack/echo/v4"
import "gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
import "gitlab.secoder.net/bauhinia/qanda-schema/ent"
import adminp "gitlab.secoder.net/bauhinia/qanda-schema/ent/admin"
import paramp "gitlab.secoder.net/bauhinia/qanda-schema/ent/param"
import "net/http"
import "golang.org/x/crypto/bcrypt"
import "encoding/hex"
import "github.com/sethvargo/go-password/password"

func Register(group *echo.Group) {
	group.POST("/login", login)
	group.POST("/add", add)
	group.GET("/list", list)
	group.GET("/param", param)
	group.POST("/param", param_edit)
	group.POST("/edit", edit)
	group.POST("/change", change)
}

// @Summary Admin Login
// @Description Login for a exsisting admin
// @Accept json
// @Produce json
// @Param body body adminLoginRequest true "admin login request"
// @Success 200 {object} adminLoginResponse "admin login response"
// @Failure 400 {string} string
// @Router /v1/admin/login [post]
func login(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(adminLoginRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	user, err := ctx.DB().Admin.Query().Where(adminp.Username(u.Username)).Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	password, err := hex.DecodeString(user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = bcrypt.CompareHashAndPassword(password, []byte(u.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	token, err := ctx.SignAdmin(user.Username)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, adminLoginResponse{
		Token: token,
		ID:    user.ID,
	})
}

type adminLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type adminLoginResponse struct {
	Token string `json:"token"`
	ID    int    `json:"id"`
}

// @Summary Admin Add
// @Description Add a new admin
// @Accept json
// @Produce json
// @Security token
// @Param body body adminAddRequest true "admin add request"
// @Success 200 {object} adminAddResponse "admin add response"
// @Failure 400 {string} string
// @Router /v1/admin/add [post]
func add(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(adminAddRequest)
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.VerifyAdmin(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	if claims.Subject != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "only admin can add admin")
	}
	pass, err := password.Generate(8, 1, 1, false, false)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := ctx.DB().Admin.Create().SetUsername(u.Username).SetPassword(hex.EncodeToString(password)).SetRole("none").Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, adminAddResponse{
		Password: pass,
		ID:       user.ID,
	})
}

type adminAddRequest struct {
	Token    string `header:"authorization" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type adminAddResponse struct {
	Password string `json:"password"`
	ID       int    `json:"id"`
}

// @Summary Admin List
// @Description List all admins
// @Produce json
// @Success 200 {object} adminListResponse "admin list response"
// @Failure 400 {string} string
// @Router /v1/admin/list [get]
func list(c echo.Context) error {
	ctx := c.(*common.Context)
	users, err := ctx.DB().Admin.Query().Order(ent.Asc(adminp.FieldID)).All(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var userlist = make([]adminInfoDisplay, 0)
	for _, user := range users {
		userlist = append(userlist, adminInfoDisplay{
			Username: user.Username,
			Role:     user.Role.String(),
		})
	}
	return ctx.JSON(http.StatusOK, adminListResponse{
		Userlist: userlist,
	})
}

type adminInfoDisplay struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type adminListResponse struct {
	Userlist []adminInfoDisplay `json:"userlist"`
}

// @Summary Param Query
// @Description Query current system param
// @Produce json
// @Security token
// @Success 200 {object} paramQueryResponse "param query response"
// @Failure 400 {string} string
// @Router /v1/admin/param [get]
func param(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(paramQueryRequest)
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	_, err := ctx.VerifyAdmin(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	pa, err := ctx.DB().Param.Query().Where(paramp.Scope("default")).Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, paramQueryResponse{
		MinPrice:       pa.MinPrice,
		MaxPrice:       pa.MaxPrice,
		AcceptDeadline: pa.AcceptDeadline,
		AnswerDeadline: pa.AnswerDeadline,
		AnswerLimit:    pa.AnswerLimit,
		DoneDeadline:   pa.DoneDeadline,
	})
}

type paramQueryRequest struct {
	Token string `header:"authorization" validate:"required"`
}

type paramQueryResponse struct {
	MinPrice       float64 `json:"min_price"`
	MaxPrice       float64 `json:"max_price"`
	AcceptDeadline int     `json:"accept_deadline"`
	AnswerDeadline int     `json:"answer_deadline"`
	AnswerLimit    int     `json:"answer_limit"`
	DoneDeadline   int     `json:"done_deadline"`
}

// @Summary Param Edit
// @Description Edit current system param
// @Accept json
// @Security token
// @Param body body paramEditRequest true "param edit request"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /v1/admin/param [post]
func param_edit(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(paramEditRequest)
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.VerifyAdmin(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	if claims.Subject != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "only admin can edit param")
	}
	upd := ctx.DB().Param.Update().Where(paramp.Scope("default"))
	if u.MinPrice != nil {
		upd = upd.SetMinPrice(*u.MinPrice)
	}
	if u.MaxPrice != nil {
		upd = upd.SetMaxPrice(*u.MaxPrice)
	}
	if u.AcceptDeadline != nil {
		upd = upd.SetAcceptDeadline(*u.AcceptDeadline)
	}
	if u.AnswerDeadline != nil {
		upd = upd.SetAnswerDeadline(*u.AnswerDeadline)
	}
	if u.AnswerLimit != nil {
		upd = upd.SetAnswerLimit(*u.AnswerLimit)
	}
	if u.DoneDeadline != nil {
		upd = upd.SetDoneDeadline(*u.DoneDeadline)
	}
	_, err = upd.Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "edit param successful")
}

type paramEditRequest struct {
	Token          string   `header:"authorization" validate:"required"`
	MinPrice       *float64 `json:"max_price"`
	MaxPrice       *float64 `json:"min_price"`
	AcceptDeadline *int     `json:"accept_deadline"`
	AnswerDeadline *int     `json:"answer_deadline"`
	AnswerLimit    *int     `json:"answer_limit"`
	DoneDeadline   *int     `json:"done_deadline"`
}

// @Summary Admin Edit
// @Description edit an administer's password
// @Accept json
// @Param body body adminEditRequest true "admin edit request"
// @Success 200 {object} string
// @Failure 400 {string} string
// @Router /v1/admin/edit [post]
func edit(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(adminEditRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.VerifyAdmin(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, err = ctx.DB().Admin.Update().Where(adminp.Username(claims.Subject)).SetPassword(hex.EncodeToString(password)).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "admin password is successfully edited")
}

type adminEditRequest struct {
	Token    string `header:"authorization" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// @Summary Admin Change
// @Description Change role of admin
// @Accept json
// @Security token
// @Param body body adminChangeRequest true "admin change request"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /v1/admin/change [post]
func change(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(adminChangeRequest)
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.VerifyAdmin(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	if claims.Subject != "admin" || u.Username == "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "only admin can change role")
	}
	_, err = ctx.DB().Admin.Update().Where(adminp.Username(u.Username)).SetRole(adminp.Role(u.Role)).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "change role successful")
}

type adminChangeRequest struct {
	Token    string `header:"authorization" validate:"required"`
	Username string `json:"username" validate:"required"`
	Role     string `json:"role", validate:"required"`
}
