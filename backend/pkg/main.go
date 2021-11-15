package qanda

import (
	"context"
	"encoding/hex"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/swaggo/echo-swagger"
	"gitlab.secoder.net/bauhinia/qanda-schema/ent"
	adminp "gitlab.secoder.net/bauhinia/qanda-schema/ent/admin"
	paramp "gitlab.secoder.net/bauhinia/qanda-schema/ent/param"
	userp "gitlab.secoder.net/bauhinia/qanda-schema/ent/user"
	questionp "gitlab.secoder.net/bauhinia/qanda-schema/ent/question"
	"gitlab.secoder.net/bauhinia/qanda/backend/pkg/admin"
	"gitlab.secoder.net/bauhinia/qanda/backend/pkg/common"
	_ "gitlab.secoder.net/bauhinia/qanda/backend/pkg/docs"
	"gitlab.secoder.net/bauhinia/qanda/backend/pkg/question"
	"gitlab.secoder.net/bauhinia/qanda/backend/pkg/user"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// @title Q&A API
// @version 1.0

// @host qanda-bauhinia.app.secoder.net
// @BasePath /

// @securityDefinitions.apikey token
// @in header
// @name Authorization
func New(serve string, storage string, database string, key string, adminKey string) *echo.Echo {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	db, err := ent.Open(storage, database)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Schema.Create(context.Background()); err != nil {
		e.Logger.Fatal(err)
	}
	c, err := db.Admin.Query().Where(adminp.Username("admin")).Count(context.Background())
	if err != nil {
		e.Logger.Fatal(err)
	}
	if c == 0 {
		password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			e.Logger.Fatal(err)
		}
		_, err = db.Admin.Create().SetUsername("admin").SetRole("admin").SetPassword(hex.EncodeToString(password)).Save(context.Background())
		if err != nil {
			e.Logger.Fatal(err)
		}
	}
	c, err = db.Param.Query().Where(paramp.Scope("default")).Count(context.Background())
	if err != nil {
		e.Logger.Fatal(err)
	}
	if c == 0 {
		_, err = db.Param.Create().SetScope("default").SetMinPrice(10).SetMaxPrice(400).SetAcceptDeadline(3600).SetAnswerDeadline(3600).SetAnswerLimit(50).SetDoneDeadline(3600).Save(context.Background())
		if err != nil {
			e.Logger.Fatal(err)
		}
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &common.Context{Context: c, DBField: db, Key: []byte(key), AdminKey: []byte(adminKey)}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", serve)
	e.GET("/docs/*", echoSwagger.WrapHandler)
	v1 := e.Group("/v1")
	user.Register(v1.Group("/user"))
	question.Register(v1.Group("/question"))
	admin.Register(v1.Group("/admin"))

	// Check question modified time EVERY SECOND 
	go func(db *ent.Client){
		pa, err := db.Param.Query().Where(paramp.Scope("default")).Only(context.Background())
		if err != nil {
			e.Logger.Fatal(err)
		}
		timeBackwardSecond := func (num int) time.Time {
			return time.Now().Add(time.Second * time.Duration(-num))
		}
		for {
			acceptClear := timeBackwardSecond(pa.AcceptDeadline)
			answerClear := timeBackwardSecond(pa.AnswerDeadline)
			doneClear := timeBackwardSecond(pa.DoneDeadline)
			// Accept Deadline
			questions, err := db.Question.Query().Where(questionp.StateEQ("reviewed")).Where(questionp.ModifiedLT(acceptClear)).WithQuestioner().All(context.Background())
			if err != nil {
				e.Logger.Fatal(err)
			}
			for _, question := range questions {
				_, err = db.Question.Update().Where(questionp.ID(question.ID)).SetState("canceled").Save(context.Background())
				if err != nil {
					e.Logger.Fatal(err)
				}
				_, err = db.User.Update().Where(userp.ID(question.Edges.Questioner.ID)).SetBalance(question.Edges.Questioner.Balance + question.Price).Save(context.Background())
				if err != nil {
					e.Logger.Fatal(err)
				}
			}
			// Answer Deadline
			questions, err = db.Question.Query().Where(questionp.StateEQ("accepted")).Where(questionp.Answered(false)).Where(questionp.ModifiedLT(answerClear)).WithQuestioner().All(context.Background())
			if err != nil {
				e.Logger.Fatal(err)
			}
			for _, question := range questions {
				_, err = db.Question.Update().Where(questionp.ID(question.ID)).SetState("canceled").Save(context.Background())
				if err != nil {
					e.Logger.Fatal(err)
				}
				_, err = db.User.Update().Where(userp.ID(question.Edges.Questioner.ID)).SetBalance(question.Edges.Questioner.Balance + question.Price).Save(context.Background())
				if err != nil {
					e.Logger.Fatal(err)
				}
			}
			// Done Deadline
			questions, err = db.Question.Query().Where(questionp.StateEQ("accepted")).Where(questionp.Answered(true)).Where(questionp.ModifiedLT(doneClear)).WithAnswerer().All(context.Background())
			if err != nil {
				e.Logger.Fatal(err)
			}
			for _, question := range questions {
				_, err = db.Question.Update().Where(questionp.ID(question.ID)).SetState("done").Save(context.Background())
				if err != nil {
					e.Logger.Fatal(err)
				}
				_, err = db.User.Update().Where(userp.ID(question.Edges.Answerer.ID)).SetBalance(question.Edges.Answerer.Balance + question.Price).Save(context.Background())
				if err != nil {
					e.Logger.Fatal(err)
				}
			}
			time.Sleep(time.Second)
		}
	}(db)

	return e
}
