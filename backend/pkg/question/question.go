package question

import "github.com/labstack/echo/v4"
import "gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
import "net/http"
import "time"

func Register(group *echo.Group) {
	group.POST("/submit", submit)
	group.GET("/:id", query)
}

// @Summary Question submit
// @Description submit a question towards the specified answerer
// @Accept json
// @Produce json
// @Param body body questionSubmitRequest true "question submit request"
// @Success 200 {object} questionSubmitResponse "question submit response"
// @Failure 400 {string} string
// @Router /v1/qestion [post]
func submit(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionSubmitRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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

// @Summary Question Query
// @Description Query a question
// @Produce json
// @Param id path string true "question query id"
// @Success 200 {object} questionQueryResponse "question query response"
// @Failure 400 {string} string
// @Router /v1/qestion/:id [get]
func query(ctx echo.Context) error {
	return echo.ErrMethodNotAllowed
}

type questionQueryResponse struct {
	AnswererId string
	Content    string
	State      string
}
