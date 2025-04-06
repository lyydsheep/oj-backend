package controller

import (
	"back/api/request"
	"back/common/app"
	"back/common/errcode"
	"back/logic/appService"
	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	svc appService.QuestionAppService
}

func NewQuestionController(svc appService.QuestionAppService) *QuestionController {
	return &QuestionController{
		svc: svc,
	}
}

func (q *QuestionController) CreateQuestion(c *gin.Context) {
	var req request.CreateRequest
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
