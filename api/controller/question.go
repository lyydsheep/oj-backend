package controller

import (
	"back/api/request"
	"back/common/app"
	"back/common/errcode"
	"back/logic/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	svc service.QuestionService
}

func NewQuestionController(svc service.QuestionService) *QuestionController {
	return &QuestionController{
		svc: svc,
	}
}

func (q *QuestionController) CreateQuestion(c *gin.Context) {
	var req request.CreateQuestion
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).Error(errcode.ErrParams.WithCause(err))
		return
	}
	// TODO 校验参数
	// TODO 补充缺省值
	// TODO 鉴权
	id, err := q.svc.CreateQuestion(c, &req)
	if err != nil {
		app.NewResponse(c).Error(err)
		return
	}
	app.NewResponse(c).Success(id)
}

func (q *QuestionController) GetQuestion(c *gin.Context) {
	questionIdStr := c.Query("question_id")
	questionId, err := strconv.ParseInt(questionIdStr, 10, 64)
	if err != nil {
		app.NewResponse(c).Error(errcode.ErrParams.WithCause(err).AppendMsg("question_id invalid, question_id: " + questionIdStr))
		return
	}
	question, err := q.svc.GetQuestionById(c, questionId)
	if err != nil {
		app.NewResponse(c).Error(err)
		return
	}
	app.NewResponse(c).Success(question)
}
