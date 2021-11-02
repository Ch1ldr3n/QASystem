package user

import "github.com/labstack/echo/v4"
import "gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
import "gitlab.secoder.net/bauhinia/qanda-schema/ent"
import userp "gitlab.secoder.net/bauhinia/qanda-schema/ent/user"
import "net/http"
import "golang.org/x/crypto/bcrypt"
import "encoding/hex"
import "strconv"

func Register(group *echo.Group) {
	group.POST("/register", register)
	group.POST("/login", login)
	group.GET("/info", info)
	group.POST("/edit", edit)
	group.GET("/filter", filter)
	group.GET("/gensig", gensig)
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := ctx.DB().User.Create().SetUsername(u.Username).SetPassword(hex.EncodeToString(password)).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = ctx.RequestTIM("im_open_login_svc", "account_import", struct {
		Identifier string `json:"Identifier"`
	}{
		Identifier: strconv.Itoa(user.ID),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	token, err := ctx.Sign(user.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, userRegisterResponse{
		Token: token,
		ID:    user.ID,
	})
}

type userRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userRegisterResponse struct {
	Token string `json:"token"`
	ID    int    `json:"id"`
}

// @Summary User Login
// @Description Login for a exsisting user
// @Accept json
// @Produce json
// @Param body body userLoginRequest true "user login request"
// @Success 200 {object} userLoginResponse "user login response"
// @Failure 400 {string} string
// @Router /v1/user/login [post]
func login(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(userLoginRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	user, err := ctx.DB().User.Query().Where(userp.Username(u.Username)).Only(ctx.Request().Context())
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
	token, err := ctx.Sign(user.Username)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, userLoginResponse{
		Token: token,
		ID:    user.ID,
	})
}

type userLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userLoginResponse struct {
	Token string `json:"token"`
	ID    int    `json:"id"`
}

// @Summary User Gensig
// @Description Gensig of current user
// @Security token
// @Produce json
// @Success 200 {object} userGensigResponse "user gensig response"
// @Failure 400 {string} string
// @Router /v1/user/gensig [get]
func gensig(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(userGensigRequest)
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	user, err := ctx.DB().User.Query().Where(userp.Username(claims.Subject)).Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	sig, err := ctx.Gensig(strconv.Itoa(user.ID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, userGensigResponse{
		Userid:    strconv.Itoa(user.ID),
		Signature: sig,
	})
}

type userGensigRequest struct {
	Token string `header:"authorization" validate:"required"`
}

type userGensigResponse struct {
	Userid    string `json:"userid"`
	Signature string `json:"signature"`
}

// @Summary User Info
// @Description Info of current user
// @Security token
// @Produce json
// @Success 200 {object} userInfoResponse "user info response"
// @Failure 400 {string} string
// @Router /v1/user/info [get]
func info(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(userInfoRequest)
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	user, err := ctx.DB().User.Query().Where(userp.Username(claims.Subject)).Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, userInfoResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		Phone:      user.Phone,
		Answerer:   user.Answerer,
		Price:      user.Price,
		Profession: user.Profession,
	})
}

type userInfoRequest struct {
	Token string `header:"authorization" validate:"required"`
}

type userInfoResponse struct {
	ID         int     `json:"id"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Answerer   bool    `json:"answerer"`
	Price      float64 `json:"price"`
	Profession string  `json:"profession"`
}

// @Summary User Edit
// @Description Edit current user
// @Security token
// @Accept json
// @Param body body userEditRequest true "user edit request"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /v1/user/edit [post]
func edit(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(userEditRequest)
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	_, err = ctx.DB().User.Update().Where(userp.Username(claims.Subject)).
		SetNillableEmail(u.Email).SetNillablePhone(u.Phone).
		SetNillableAnswerer(u.Answerer).SetNillablePrice(u.Price).
		SetNillableProfession(u.Profession).
		Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.String(http.StatusOK, "user info updated")
}

type userEditRequest struct {
	Token      string   `header:"authorization" validate:"required"`
	Email      *string  `json:"email"`
	Phone      *string  `json:"phone"`
	Answerer   *bool    `json:"answerer"`
	Price      *float64 `json:"price"`
	Profession *string  `json:"profession"`
}

// @Summary User Filter
// @Description Filter for wanted users
// @Produce json
// @Param query query userFilterRequest true "user filter request"
// @Success 200 {object} userFilterResponse "user filter response"
// @Failure 400 {string} string
// @Router /v1/user/filter [get]
func filter(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(userFilterRequest)
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	// begin Filter
	users := ctx.DB().User.Query().Where(userp.UsernameNEQ(claims.Subject))
	if u.ID != nil {
		users = users.Where(userp.ID(*u.ID))
	}
	if u.Username != nil {
		users = users.Where(userp.UsernameContains(*u.Username))
	}
	if u.Email != nil {
		users = users.Where(userp.EmailContains(*u.Email))
	}
	if u.Phone != nil {
		users = users.Where(userp.Phone(*u.Phone))
	}
	if u.Answerer != nil {
		users = users.Where(userp.Answerer(*u.Answerer))
	}
	if u.PriceLowerBound != nil {
		users = users.Where(userp.PriceGT(*u.PriceLowerBound))
	}
	if u.PriceUpperBound != nil {
		users = users.Where(userp.PriceLT(*u.PriceUpperBound))
	}
	if u.Profession != nil {
		users = users.Where(userp.Profession(*u.Profession))
	}
	// print result
	const numLimit = 1000
	candidates, err := users.Order(ent.Asc(userp.FieldID)).Limit(numLimit).All(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var userlist [numLimit]userInfoDisplay
	listlen := len(candidates)
	for i := 0; i < listlen; i++ {
		userlist[i].ID = candidates[i].ID
		userlist[i].Username = candidates[i].Username
		userlist[i].Email = candidates[i].Email
		userlist[i].Phone = candidates[i].Phone
		userlist[i].Answerer = candidates[i].Answerer
		userlist[i].Price = candidates[i].Price
		userlist[i].Profession = candidates[i].Profession
	}
	return ctx.JSON(http.StatusOK, userFilterResponse{
		ResultNum: listlen,
		Userlist:  userlist[:listlen],
	})
}

type userFilterRequest struct {
	Token           string   `header:"authorization" validate:"required"`
	ID              *int     `query:"id"`
	Username        *string  `query:"username"`
	Email           *string  `query:"email"`
	Phone           *string  `query:"phone"`
	Answerer        *bool    `query:"answerer"`
	PriceUpperBound *float64 `query:"priceUpperBound"`
	PriceLowerBound *float64 `query:"priceLowerBound"`
	Profession      *string  `query:"profession"`
}

type userInfoDisplay = userInfoResponse

type userFilterResponse struct {
	ResultNum int               `json:"num"`
	Userlist  []userInfoDisplay `json:"userlist"`
}
