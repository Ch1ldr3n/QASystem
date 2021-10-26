package question

import "github.com/labstack/echo/v4"
import "gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
import userp "gitlab.secoder.net/bauhinia/qanda-schema/ent/user"
import questionp "gitlab.secoder.net/bauhinia/qanda-schema/ent/question"
import "net/http"
import "time"


func Register(group *echo.Group) {
	group.POST("/submit", submit)
	group.POST("/pay", pay)
	group.GET("/:id", query)
}

// @Summary Question Submit
// @Description Submit a question towards the specified answerer
// @Accept json
// @Produce json
// @Param body body questionSubmitRequest true "question submit request"
// @Success 200 {object} questionSubmitResponse "question submit response"
// @Failure 400 {string} string
// @Router /v1/question/summit [post]
func submit(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionSubmitRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if u.QuestionerID == u.AnswererID {
		return echo.NewHTTPError(http.StatusBadRequest, "error: questioner and answerer being the same person")
	}
	question, err := ctx.DB().Question.Create().
		SetPrice(u.Price).
		SetTitle(u.Title).
		SetContent(u.Content).
		SetCreated(time.Now()).
		SetState("created").
		SetQuestionerID(u.QuestionerID).
		SetAnswererID(u.AnswererID).
		Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, questionSubmitResponse{
		QuestionId: question.ID,
	})
}

type questionSubmitRequest struct {
	Price      float64	`json:"price"`
	Title	   string	`json:"title"`
	Content    string	`json:"content"`
	QuestionerID	int	`json:"questionerid"`
	AnswererID	int	`json:"answererid"`
}

type questionSubmitResponse struct {
	QuestionId int	`json:"questionid"`
}

// @Summary Question Pay
// @Description Pay for a question
// @Accept json
// @Param body body questionPayRequest true "question pay request"
// @Success 200 {object} string "question pay response"
// @Failure 400 {string} string
// @Router /v1/question/pay [post]
func pay(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionPayRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	question, err := ctx.DB().Question.Query().Where(questionp.ID(u.QuestionID)).WithQuestioner().WithAnswerer().Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if question.State != "created" {
		return echo.NewHTTPError(http.StatusBadRequest, "question state is not 'created'")
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	payer, payee := question.Edges.Questioner, question.Edges.Answerer
	if claims.Subject != payer.Username {
		return echo.NewHTTPError(http.StatusBadRequest, "current user is not the questioner")
	}
	if payer.Balance < question.Price {
		return echo.NewHTTPError(http.StatusBadRequest, "payer's balance is less than question price")
	}
	_, err = ctx.DB().User.Update().Where(userp.ID(payer.ID)).SetBalance(payer.Balance + question.Price).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, err = ctx.DB().User.Update().Where(userp.ID(payee.ID)).SetBalance(payee.Balance + question.Price).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState("paid").Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "payment succeeded")
}

type questionPayRequest struct {
	QuestionID	int	`json:"questionid"`
	Token	string   `header:"authorization" validate:"required"`
}

// @Summary Question Query
// @Description Query a question
// @Produce json
// @Param id path string true "question query id"
// @Success 200 {object} questionQueryResponse "question query response"
// @Failure 400 {string} string
// @Router /v1/question/:id [get]
func query(ctx echo.Context) error {
	return echo.ErrMethodNotAllowed
}

type questionQueryResponse struct {
	AnswererId string
	Content    string
	State      string
}
