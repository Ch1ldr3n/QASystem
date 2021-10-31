package question

import "github.com/labstack/echo/v4"
import "gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
import userp "gitlab.secoder.net/bauhinia/qanda-schema/ent/user"
import questionp "gitlab.secoder.net/bauhinia/qanda-schema/ent/question"
import "net/http"
import "time"
import "strconv"

func Register(group *echo.Group) {
	group.POST("/submit", submit)
	group.POST("/pay", pay)
	group.GET("/:id", query)
	group.GET("/list", list)
	group.GET("/mine", mine)
	group.POST("/accept", accept)
	group.POST("/close", close)
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
	payer, _ := question.Edges.Questioner, question.Edges.Answerer
	if claims.Subject != payer.Username {
		return echo.NewHTTPError(http.StatusBadRequest, "current user is not the questioner")
	}
	if payer.Balance < question.Price {
		return echo.NewHTTPError(http.StatusBadRequest, "payer's balance is less than question price")
	}
	_, err = ctx.DB().User.Update().Where(userp.ID(payer.ID)).SetBalance(payer.Balance - question.Price).Save(ctx.Request().Context())
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
func query(c echo.Context) error {
	ctx := c.(*common.Context)
	idstring := c.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	question, err1 := ctx.DB().Question.Query().Where(questionp.ID(id)).Only(ctx.Request().Context())
	if err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	return ctx.JSON(http.StatusOK, questionQueryResponse{
		Price:	question.Price,
		Title:	question.Title,
		Content:	question.Content,
		State:	string(question.State),
	})
}

type questionQueryResponse struct {
	Price      float64	`json:"price"`
	Title	   string	`json:"title"`
	Content    string	`json:"content"`
	State	   string	`json:"state"`
}

// @Summary Question List
// @Description List of all questions open to all users
// @Produce json
// @Success 200 {object} questionListResponse "question list response"
// @Failure 400 {string} string
// @Router /v1/question/list [get]
func list(c echo.Context) error {
	ctx := c.(*common.Context)
	const numLimit = 1000
	var questionlist [numLimit]questionInfoDesplay
	questions, err := ctx.DB().Question.Query().Limit(numLimit).WithQuestioner().WithAnswerer().All(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	listlen := len(questions)
	for i := 0; i < listlen; i = i + 1 {
		questionlist[i].Price = questions[i].Price;
		questionlist[i].Title = questions[i].Title;
		questionlist[i].Content = questions[i].Content;
		questionlist[i].State = string(questions[i].State);
		questionlist[i].QuestionerID = questions[i].Edges.Questioner.ID;
		questionlist[i].AnswererID = questions[i].Edges.Answerer.ID;
	}
	return ctx.JSON(http.StatusOK, questionListResponse{
		ResultNum:	listlen,
		Questionlist:	questionlist[:listlen],
	})
}

type questionInfoDesplay struct {
	Price      float64	`json:"price"`
	Title	   string	`json:"title"`
	Content    string	`json:"content"`
	State	   string	`json:"state"`
	QuestionerID	int	`json:"questionerid"`
	AnswererID	int	`json:"answererid"`
}

type questionListResponse struct {
	ResultNum	int	`json:"num"`
	Questionlist	[]questionInfoDesplay	`json:"questionlist"`
}

// @Summary Question Mine
// @Description List of all relevant questions
// @Accept json
// @Produce json
// @Param body body questionMineRequest true "question mine request"
// @Success 200 {object} questionMineResponse "question mine response"
// @Failure 400 {string} string
// @Router /v1/question/mine [get]
func mine(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionMineRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
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
	const numLimit = 1000
	var askedlist [numLimit]questionInfoDesplay
	var answeredlist [numLimit]questionInfoDesplay
	user, err1 := ctx.DB().User.Query().Where(userp.Username(claims.Subject)).WithAsked().WithAnswered().Only(ctx.Request().Context())
	if err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	// asked
	listlen1 := len(user.Edges.Asked)
	for i := 0; i < listlen1; i = i + 1 {
		askedlist[i].Price = user.Edges.Asked[i].Price;
		askedlist[i].Title = user.Edges.Asked[i].Title;
		askedlist[i].Content = user.Edges.Asked[i].Content;
		askedlist[i].State = string(user.Edges.Asked[i].State);
		// get its answerer
		question, err := ctx.DB().Question.Query().Where(questionp.ID(user.Edges.Asked[i].ID)).WithAnswerer().Only(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		askedlist[i].QuestionerID = user.ID;
		askedlist[i].AnswererID = question.Edges.Answerer.ID;
	}
	// answered
	listlen2 := len(user.Edges.Answered)
	for i := 0; i < listlen2; i = i + 1 {
		answeredlist[i].Price = user.Edges.Answered[i].Price;
		answeredlist[i].Title = user.Edges.Answered[i].Title;
		answeredlist[i].Content = user.Edges.Answered[i].Content;
		answeredlist[i].State = string(user.Edges.Answered[i].State);
		//get its questioner
		question, err := ctx.DB().Question.Query().Where(questionp.ID(user.Edges.Answered[i].ID)).WithQuestioner().Only(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		answeredlist[i].QuestionerID = question.Edges.Questioner.ID;
		answeredlist[i].AnswererID = user.ID;
	}
	return ctx.JSON(http.StatusOK, questionMineResponse{
		AskedNum: listlen1,
		AskedList: askedlist[:listlen1],
		AnsweredNum: listlen2, 
		AnsweredList: answeredlist[:listlen2], 
	})
}

type questionMineRequest struct {
	Token	string   `header:"authorization" validate:"required"`
}

type questionMineResponse struct {
	AskedNum	int	`json:"askednum"`
	AskedList	[]questionInfoDesplay	`json:"askedlist"`
	AnsweredNum	int	`json:"answerednum"`
	AnsweredList	[]questionInfoDesplay	`json:"answeredlist"`
}

// @Summary Question Accept
// @Description Accept a question
// @Accept json
// @Param body body questionAcceptRequest true "question accept request"
// @Success 200 {object} string "question accept response"
// @Failure 400 {string} string
// @Router /v1/question/accept [post]
func accept(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionAcceptRequest)
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
	if question.State != "paid" {
		return echo.NewHTTPError(http.StatusBadRequest, "question state is not 'paid'")
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	questioner, answerer := question.Edges.Questioner, question.Edges.Answerer
	if claims.Subject != answerer.Username {
		return echo.NewHTTPError(http.StatusBadRequest, "current user is not the answerer")
	}
	if u.Choice {
		_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState("accepted").Save(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return ctx.JSON(http.StatusOK, "question is accepted")
	} else {
		_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState("canceled").Save(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		_, err = ctx.DB().User.Update().Where(userp.ID(questioner.ID)).SetBalance(questioner.Balance + question.Price).Save(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return ctx.JSON(http.StatusOK, "question is canceled")
	}
}

type questionAcceptRequest struct {
	QuestionID	int	`json:"questionid"`
	Choice	bool	`json:"choice"`
	Token	string   `header:"authorization" validate:"required"`
}

// @Summary Question Close
// @Description Close a question; Questioner only
// @Accept json
// @Param body body questionCloseRequest true "question close request"
// @Success 200 {object} string "question close response"
// @Failure 400 {string} string
// @Router /v1/question/close [post]
func close(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionCloseRequest)
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
	if question.State != "accepted" {
		return echo.NewHTTPError(http.StatusBadRequest, "question state is not 'accepted'")
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	questioner, answerer := question.Edges.Questioner, question.Edges.Answerer
	if claims.Subject != questioner.Username {
		return echo.NewHTTPError(http.StatusBadRequest, "current user is not the questioner")
	}
	_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState("done").Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, err = ctx.DB().User.Update().Where(userp.ID(answerer.ID)).SetBalance(answerer.Balance + question.Price).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "question is done")
}

type questionCloseRequest struct {
	QuestionID	int	`json:"questionid"`
	Token	string   `header:"authorization" validate:"required"`
}