package question

import "github.com/labstack/echo/v4"

func Register(group *echo.Group) {
	group.POST("/", create)
}

// @Summary Question Create
// @Description Create a question towards the specified answerer
// @Accept json
// @Produce json
// @Param body body questionCreateRequest true "question create request"
// @Success 200 {object} questionCreateResponse "question create response"
// @Failure 400 {string} string
// @Router /v1/qestion [post]
func create(ctx echo.Context) error {
	return echo.ErrMethodNotAllowed
}

type questionCreateRequest struct {
	AnswererId string
	Content    string
}

type questionCreateResponse struct {
	QuestionId string
}
