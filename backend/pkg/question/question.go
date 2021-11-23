package question

import "github.com/labstack/echo/v4"
import "gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
import "gitlab.secoder.net/bauhinia/qanda-schema/ent"
import userp "gitlab.secoder.net/bauhinia/qanda-schema/ent/user"
import adminp "gitlab.secoder.net/bauhinia/qanda-schema/ent/admin"
import questionp "gitlab.secoder.net/bauhinia/qanda-schema/ent/question"
import "net/http"
import "time"
import "strconv"
import "sort"

func Register(group *echo.Group) {
	group.POST("/submit", submit)
	group.POST("/pay", pay)
	group.GET("/:id", query)
	group.GET("/list", list)
	group.GET("/mine", mine)
	group.POST("/accept", accept)
	group.POST("/review", review)
	group.POST("/close", close)
	group.POST("/cancel", cancel)
	group.GET("/review", revlist)
	group.GET("/aggreg", aggreg)
	group.POST("/callback", callback)
}

// @Summary Question Submit
// @Description Submit a question towards the specified answerer
// @Security token
// @Accept json
// @Produce json
// @Param body body questionSubmitRequest true "question submit request"
// @Success 200 {object} questionSubmitResponse "question submit response"
// @Failure 400 {string} string
// @Router /v1/question/submit [post]
func submit(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionSubmitRequest)
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
	questioner, err := ctx.DB().User.Query().Where(userp.Username(claims.Subject)).Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if questioner.ID == u.AnswererID {
		return echo.NewHTTPError(http.StatusBadRequest, "error: questioner and answerer being the same person")
	}
	answerer, err0 := ctx.DB().User.Query().Where(userp.ID(u.AnswererID)).Only(ctx.Request().Context())
	if err0 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err0.Error())
	}
	if !answerer.Answerer {
		return echo.NewHTTPError(http.StatusBadRequest, "error: asking those who are not qualified as an answerer")
	}
	question, err := ctx.DB().Question.Create().
		SetPrice(answerer.Price).
		SetTitle(u.Title).
		SetContent(u.Content).
		SetCreated(time.Now()).
		SetModified(time.Now()).
		SetState(questionp.StateCreated).
		SetQuestionerID(questioner.ID).
		SetAnswererID(answerer.ID).
		SetMsgCount(0).
		SetAnswered(false).
		SetPublic(true).
		Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = ctx.RequestTIM("group_open_http_svc", "create_group", struct {
		Type       string `json:"Type"`
		GroupId    string `json:"GroupId"`
		Name       string `json:"Name"`
		MemberList []struct {
			Member_Account string `json:"Member_Account"`
		} `json:"MemberList"`
	}{
		Type:    "Private",
		GroupId: strconv.Itoa(question.ID),
		Name:    "qanda",
		MemberList: []struct {
			Member_Account string `json:"Member_Account"`
		}{{Member_Account: strconv.Itoa(questioner.ID)}, {Member_Account: strconv.Itoa(answerer.ID)}, {Member_Account: "public"}},
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, questionSubmitResponse{
		QuestionId: question.ID,
	})
}

type questionSubmitRequest struct {
	Token      string `header:"authorization" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	AnswererID int    `json:"answererid" validate:"required"`
}

type questionSubmitResponse struct {
	QuestionId int `json:"questionid"`
}

// @Summary Question Pay
// @Description Pay for a question
// @Security token
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
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	question, err := ctx.DB().Question.Query().Where(questionp.ID(u.QuestionID)).WithQuestioner().WithAnswerer().Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if question.State != questionp.StateCreated {
		return echo.NewHTTPError(http.StatusBadRequest, "question state is not 'created'")
	}
	payer, _ := question.Edges.Questioner, question.Edges.Answerer
	if claims.Subject != payer.Username {
		return echo.NewHTTPError(http.StatusBadRequest, "current user is not the questioner")
	}
	if payer.Balance < question.Price {
		payer.Balance = question.Price
	}
	_, err = ctx.DB().User.Update().Where(userp.ID(payer.ID)).SetBalance(payer.Balance - question.Price).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState(questionp.StatePaid).SetModified(time.Now()).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "payment succeeded")
}

type questionPayRequest struct {
	QuestionID int    `json:"questionid"`
	Token      string `header:"authorization" validate:"required"`
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
	question, err1 := ctx.DB().Question.Query().Where(questionp.ID(id)).WithQuestioner().WithAnswerer().Only(ctx.Request().Context())
	if err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	return ctx.JSON(http.StatusOK, questionQueryResponse{
		Price:              question.Price,
		Title:              question.Title,
		Content:            question.Content,
		State:              string(question.State),
		QuestionerID:       question.Edges.Questioner.ID,
		AnswererID:         question.Edges.Answerer.ID,
		QuestionerUsername: question.Edges.Questioner.Username,
		AnswererUsername:   question.Edges.Answerer.Username,
	})
}

type questionQueryResponse = questionInfoDisplay

// @Summary Question List
// @Description List of all done and public questions open to all users
// @Produce json
// @Success 200 {object} questionListResponse "question list response"
// @Failure 400 {string} string
// @Router /v1/question/list [get]
func list(c echo.Context) error {
	ctx := c.(*common.Context)
	const numLimit = 1000
	var questionlist [numLimit]questionInfoDisplay
	questions, err := ctx.DB().Question.Query().
		Where(questionp.StateEQ(questionp.StateDone)).
		Where(questionp.Public(true)).
		Order(ent.Desc(questionp.FieldID)).Limit(numLimit).
		WithQuestioner().WithAnswerer().
		All(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	listlen := len(questions)
	for i := 0; i < listlen; i = i + 1 {
		questionlist[i].ID = questions[i].ID
		questionlist[i].Price = questions[i].Price
		questionlist[i].Title = questions[i].Title
		questionlist[i].Content = questions[i].Content
		questionlist[i].State = string(questions[i].State)
		questionlist[i].QuestionerID = questions[i].Edges.Questioner.ID
		questionlist[i].AnswererID = questions[i].Edges.Answerer.ID
		questionlist[i].QuestionerUsername = questions[i].Edges.Questioner.Username
		questionlist[i].AnswererUsername = questions[i].Edges.Answerer.Username
	}
	return ctx.JSON(http.StatusOK, questionListResponse{
		ResultNum:    listlen,
		Questionlist: questionlist[:listlen],
	})
}

type questionInfoDisplay struct {
	ID                 int     `json:"id"`
	Price              float64 `json:"price"`
	Title              string  `json:"title"`
	Content            string  `json:"content"`
	State              string  `json:"state"`
	QuestionerID       int     `json:"questionerid"`
	AnswererID         int     `json:"answererid"`
	QuestionerUsername string  `json:"qusername"`
	AnswererUsername   string  `json:"ausername"`
}

type questionListResponse struct {
	ResultNum    int                   `json:"num"`
	Questionlist []questionInfoDisplay `json:"questionlist"`
}

// @Summary Question Mine
// @Description List of all relevant questions
// @Security token
// @Produce json
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
	var askedlist [numLimit]questionInfoDisplay
	var answeredlist [numLimit]questionInfoDisplay
	user, err1 := ctx.DB().User.Query().Where(userp.Username(claims.Subject)).
		WithAsked(func(q *ent.QuestionQuery) {
			q.Order(ent.Desc(questionp.FieldID))
			q.WithAnswerer()
		}).
		WithAnswered(func(q *ent.QuestionQuery) {
			q.Where(questionp.StateNotIn(questionp.StateCreated, questionp.StatePaid))
			q.Order(ent.Desc(questionp.FieldID))
			q.WithQuestioner()
		}).
		Only(ctx.Request().Context())
	if err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	// asked
	listlen1 := len(user.Edges.Asked)
	for i := 0; i < listlen1; i = i + 1 {
		askedlist[i].ID = user.Edges.Asked[i].ID
		askedlist[i].Price = user.Edges.Asked[i].Price
		askedlist[i].Title = user.Edges.Asked[i].Title
		askedlist[i].Content = user.Edges.Asked[i].Content
		askedlist[i].State = string(user.Edges.Asked[i].State)
		askedlist[i].QuestionerID = user.ID
		askedlist[i].AnswererID = user.Edges.Asked[i].Edges.Answerer.ID
		askedlist[i].QuestionerUsername = user.Username
		askedlist[i].AnswererUsername = user.Edges.Asked[i].Edges.Answerer.Username
	}
	// answered
	listlen2 := len(user.Edges.Answered)
	for i := 0; i < listlen2; i = i + 1 {
		answeredlist[i].ID = user.Edges.Answered[i].ID
		answeredlist[i].Price = user.Edges.Answered[i].Price
		answeredlist[i].Title = user.Edges.Answered[i].Title
		answeredlist[i].Content = user.Edges.Answered[i].Content
		answeredlist[i].State = string(user.Edges.Answered[i].State)
		answeredlist[i].QuestionerID = user.Edges.Answered[i].Edges.Questioner.ID
		answeredlist[i].AnswererID = user.ID
		answeredlist[i].QuestionerUsername = user.Edges.Answered[i].Edges.Questioner.Username
		answeredlist[i].AnswererUsername = user.Username
	}
	return ctx.JSON(http.StatusOK, questionMineResponse{
		AskedNum:     listlen1,
		AskedList:    askedlist[:listlen1],
		AnsweredNum:  listlen2,
		AnsweredList: answeredlist[:listlen2],
	})
}

type questionMineRequest struct {
	Token string `header:"authorization" validate:"required"`
}

type questionMineResponse struct {
	AskedNum     int                   `json:"askednum"`
	AskedList    []questionInfoDisplay `json:"askedlist"`
	AnsweredNum  int                   `json:"answerednum"`
	AnsweredList []questionInfoDisplay `json:"answeredlist"`
}

// @Summary Question Accept
// @Description Accept a question
// @Security token
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
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	question, err := ctx.DB().Question.Query().Where(questionp.ID(u.QuestionID)).WithQuestioner().WithAnswerer().Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if question.State != questionp.StateReviewed {
		return echo.NewHTTPError(http.StatusBadRequest, "question state is not 'reviewed'")
	}
	questioner, answerer := question.Edges.Questioner, question.Edges.Answerer
	if claims.Subject != answerer.Username {
		return echo.NewHTTPError(http.StatusBadRequest, "current user is not the answerer")
	}
	if u.Choice {
		_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState(questionp.StateAccepted).SetModified(time.Now()).Save(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return ctx.JSON(http.StatusOK, "question is accepted")
	} else {
		_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState(questionp.StateCanceled).SetModified(time.Now()).Save(ctx.Request().Context())
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
	QuestionID int    `json:"questionid"`
	Choice     bool   `json:"choice"`
	Token      string `header:"authorization" validate:"required"`
}

// @Summary Question Review
// @Description Review a question
// @Security token
// @Accept json
// @Param body body questionReviewRequest true "question review request"
// @Success 200 {object} string "question review response"
// @Failure 400 {string} string
// @Router /v1/question/review [post]
func review(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionReviewRequest)
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
	question, err := ctx.DB().Question.Query().Where(questionp.ID(u.QuestionID)).WithQuestioner().WithAnswerer().Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if question.State != questionp.StatePaid {
		return echo.NewHTTPError(http.StatusBadRequest, "question state is not 'paid'")
	}
	admin, err := ctx.DB().Admin.Query().Where(adminp.Username(claims.Subject)).Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if admin.Role == "none" {
		return echo.NewHTTPError(http.StatusForbidden, "admin with non role cannot review questions")
	}
	questioner := question.Edges.Questioner
	if u.Choice {
		_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState(questionp.StateReviewed).SetModified(time.Now()).Save(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return ctx.JSON(http.StatusOK, "question is reviewed")
	} else {
		_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState(questionp.StateCanceled).SetModified(time.Now()).Save(ctx.Request().Context())
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

type questionReviewRequest struct {
	QuestionID int    `json:"questionid"`
	Choice     bool   `json:"choice"`
	Token      string `header:"authorization" validate:"required"`
}

// @Summary Question Close
// @Description Close a question; Questioner only
// @Security token
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
	if err := ctx.Validate(u); err != nil {
		return err
	}
	claims, err := ctx.Verify(u.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}
	question, err := ctx.DB().Question.Query().Where(questionp.ID(u.QuestionID)).WithQuestioner().WithAnswerer().Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if question.State != questionp.StateAccepted {
		return echo.NewHTTPError(http.StatusBadRequest, "question state is not 'accepted'")
	}
	questioner, answerer := question.Edges.Questioner, question.Edges.Answerer
	if claims.Subject != questioner.Username && claims.Subject != answerer.Username {
		return echo.NewHTTPError(http.StatusBadRequest, "current user is neither the questioner nor the answerer")
	}
	_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState(questionp.StateDone).SetModified(time.Now()).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, err = ctx.DB().User.Update().Where(userp.ID(answerer.ID)).SetBalance(answerer.Balance + question.Price).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Send a system notice
	err = ctx.RequestTIM("group_open_http_svc", "send_group_system_notification", struct {
		GroupId string `json:"GroupId"`
		Content string `json:"Content"`
	}{
		GroupId: strconv.Itoa(question.ID),
		Content: "--- Question is Done ---",
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "question is done")
}

type questionCloseRequest struct {
	QuestionID int    `json:"questionid"`
	Token      string `header:"authorization" validate:"required"`
}

// @Summary Question Cancel
// @Description Cancel a question; Questioner or answerer only
// @Security token
// @Accept json
// @Param body body questionCancelRequest true "question cancel request"
// @Success 200 {object} string "question cancel response"
// @Failure 400 {string} string
// @Router /v1/question/cancel [post]
func cancel(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionCancelRequest)
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
	question, err := ctx.DB().Question.Query().Where(questionp.ID(u.QuestionID)).WithQuestioner().WithAnswerer().Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	questioner, answerer := question.Edges.Questioner, question.Edges.Answerer
	if claims.Subject != questioner.Username && claims.Subject != answerer.Username {
		return echo.NewHTTPError(http.StatusBadRequest, "current user is neither the questioner nor the answerer")
	}
	if question.State == questionp.StatePaid || question.State == questionp.StateReviewed || question.State == questionp.StateAccepted {
		_, err = ctx.DB().User.Update().Where(userp.ID(questioner.ID)).SetBalance(questioner.Balance + question.Price).Save(ctx.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	} else if question.State != questionp.StateCreated {
		return echo.NewHTTPError(http.StatusBadRequest, "question state is not 'created', 'paid', 'reviewed' or 'accepted'")
	}
	_, err = ctx.DB().Question.Update().Where(questionp.ID(question.ID)).SetState(questionp.StateCanceled).SetModified(time.Now()).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "question is canceled")
}

type questionCancelRequest struct {
	QuestionID int    `json:"questionid"`
	Token      string `header:"authorization" validate:"required"`
}

// @Summary Review List
// @Description List all to-be-reviewed questions
// @Produce json
// @Security token
// @Success 200 {object} questionRevlistResponse "question review list response"
// @Failure 400 {string} string
// @Router /v1/question/review [get]
func revlist(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionRevlistRequest)
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
	questions, err := ctx.DB().Question.Query().Where(questionp.StateEQ(questionp.StatePaid)).WithQuestioner().WithAnswerer().Order(ent.Desc(questionp.FieldID)).All(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var questionlist = make([]questionInfoDisplay, 0)
	for _, question := range questions {
		questionlist = append(questionlist, questionInfoDisplay{
			ID:                 question.ID,
			Price:              question.Price,
			Title:              question.Title,
			Content:            question.Content,
			State:              question.State.String(),
			QuestionerID:       question.Edges.Questioner.ID,
			AnswererID:         question.Edges.Answerer.ID,
			QuestionerUsername: question.Edges.Questioner.Username,
			AnswererUsername:   question.Edges.Answerer.Username,
		})
	}
	return ctx.JSON(http.StatusOK, questionRevlistResponse{
		Number:       len(questionlist),
		QuestionList: questionlist,
	})
}

type questionRevlistRequest struct {
	Token string `header:"authorization" validate:"required"`
}

type questionRevlistResponse struct {
	Number       int                   `json:"number"`
	QuestionList []questionInfoDisplay `json:"questionlist"`
}

// @Summary Question Aggreg
// @Description Summary of income of this month
// @Security token
// @Produce json
// @Success 200 {object} questionAggregResponse "question aggreg response"
// @Failure 400 {string} string
// @Router /v1/question/aggreg [get]
func aggreg(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionAggregRequest)
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
	// question filter:
	user, err1 := ctx.DB().User.Query().Where(userp.Username(claims.Subject)).
		WithAsked(func(q *ent.QuestionQuery) {
			q.Where(questionp.StateEQ(questionp.StateDone))
			q.WithQuestioner()
			q.WithAnswerer()
		}).
		WithAnswered(func(q *ent.QuestionQuery) {
			q.Where(questionp.StateEQ(questionp.StateDone))
			q.WithQuestioner()
			q.WithAnswerer()
		}).
		Only(ctx.Request().Context())
	if err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	// calculation:
	data := make(map[int](map[int](*questionAggregMonthData)))
	// - Asked
	for _, question := range append(user.Edges.Asked, user.Edges.Answered...) {
		timeObj := time.Unix(question.Modified.Unix(), 0)
		year, month := int(timeObj.Year()), int(timeObj.Month())
		// insert
		if _, exist := data[year]; !exist {
			data[year] = make(map[int](*questionAggregMonthData))
		}
		if _, exist := data[year][month]; !exist {
			data[year][month] = new(questionAggregMonthData)
			data[year][month].Year = year
			data[year][month].Month = month
		}
		// update
		if question.Edges.Answerer.ID == user.ID {
			data[year][month].Earning += question.Price
		}
		if question.Edges.Questioner.ID == user.ID {
			data[year][month].Spending += question.Price
		}
	}
	list := make([]questionAggregMonthData, 0)
	for year := range data {
		for month := range data[year] {
			list = append(list, *data[year][month])
		}
	}
	li := questionAggregMonthDataList(list)
	sort.Sort(li)
	return ctx.JSON(http.StatusOK, questionAggregResponse{
		Data: li,
	})
}

type questionAggregRequest struct {
	Token string `header:"authorization" validate:"required"`
}

type questionAggregMonthData struct {
	Year     int
	Month    int
	Earning  float64 // default: 0.0
	Spending float64 // default: 0.0
}

type questionAggregMonthDataList []questionAggregMonthData

func (md questionAggregMonthDataList) Len() int {
	return len(md)
}
func (md questionAggregMonthDataList) Less(i, j int) bool {
	return md[i].Year < md[j].Year || (md[i].Year == md[j].Year && md[i].Month < md[j].Month)
}
func (md questionAggregMonthDataList) Swap(i, j int) {
	md[i], md[j] = md[j], md[i]
}

type questionAggregResponse struct {
	Data []questionAggregMonthData `json:"list"`
}

// @Summary Question Callback
// @Description Response to IM callback
// @Accept json
// @Produce json
// @Param body body questionCallbackRequest true "question callback request"
// @Success 200 {object} string "question callback response"
// @Failure 400 {string} string
// @Router /v1/question/callback [post]
func callback(c echo.Context) error {
	ctx := c.(*common.Context)
	u := new(questionCallbackRequest)
	if err := ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}
	groupid, err := strconv.Atoi(u.GroupId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	senderid, err := strconv.Atoi(u.From_Account)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	question, err := ctx.DB().Question.Query().Where(questionp.ID(groupid)).WithAnswerer().Only(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	upd := ctx.DB().Question.Update().Where(questionp.ID(groupid)).SetMsgCount(question.MsgCount + 1).SetModified(time.Now())
	if !question.Answered && senderid == question.Edges.Answerer.ID {
		upd = upd.SetAnswered(true)
	}
	if _, err := upd.Save(ctx.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, questionCallbackResponse{
		ActionStatus: "OK",
		ErrorCode:    0,
		ErrorInfo:    "",
	})
}

type questionCallbackRequest struct {
	SdkAppid        string `query:"SdkAppid"`
	CallbackCommand string `json:"CallbackCommand"`
	GroupId         string `json:"GroupId"`
	Type            string `json:"Type"`
	From_Account    string `json:"From_Account"`
	MsgSeq          int    `json:"MsgSeq"`
	MsgTime         int    `json:"MsgTime"`
}

type questionCallbackResponse struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
}
