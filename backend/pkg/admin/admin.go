package admin

import "github.com/labstack/echo/v4"
import "gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
import adminp "gitlab.secoder.net/bauhinia/qanda-schema/ent/admin"
import "net/http"
import "golang.org/x/crypto/bcrypt"
import "encoding/hex"
import "github.com/sethvargo/go-password/password"

func Register(group *echo.Group) {
	group.POST("/login", login)
	group.POST("/add", add)
	group.GET("/list", list)
	// group.POST("/edit", edit)
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
	users, err := ctx.DB().Admin.Query().All(ctx.Request().Context())
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
