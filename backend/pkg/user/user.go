package user

import "github.com/labstack/echo/v4"
import "gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
import "net/http"
import "golang.org/x/crypto/bcrypt"
import "encoding/hex"

func Register(group *echo.Group) {
	group.POST("/register", register)
	group.POST("/login", login)
}

// @Summary User Register
// @Description Register a new user
// @Accept json
// @Produce json
// @Param body body userRegisterRequest true "user register request"
// @Success 200 {object} userRegisterResponse "user register response"
// @Failure 400 {string} string
// @Router /v1/user/register [post]
func register(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(userRegisterRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = ctx.DB().User.Create().SetUsername(u.Username).SetPassword(hex.EncodeToString(password)).Save(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, userRegisterResponse{
		Token: "example",
	})
}

type userRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userRegisterResponse struct {
	Token string `json:"token"`
}

// @Summary User Login
// @Description Login for a exsisting user
// @Accept json
// @Produce json
// @Param body body userLoginRequest true "user login request"
// @Success 200 {object} userLoginRequest "user login request"
// @Failure 400 {string} string
// @Router /v1/user/login [post]
func login(ctx echo.Context) error {
	u := new(userLoginRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, userLoginResponse{
		Token: "example",
	})
}

type userLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userLoginResponse struct {
	Token string `json:"token"`
}
