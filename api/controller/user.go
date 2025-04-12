package controller

import (
	"back/api/request"
	"back/common/app"
	"back/common/errcode"
	"back/logic/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc service.UserService
}

func NewUserController(svc service.UserService) *UserController {
	return &UserController{svc: svc}
}

func (c *UserController) Login(ctx *gin.Context) {
	var user request.User
	if err := ctx.ShouldBind(&user); err != nil {
		app.NewResponse(ctx).Error(errcode.ErrParams.WithCause(err).AppendMsg("参数错误"))
		return
	}
	if err := c.svc.Login(ctx, &user); err != nil {
		app.NewResponse(ctx).Error(err)
		return
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var user request.User
	var err error
	if err = ctx.ShouldBind(&user); err != nil {
		app.NewResponse(ctx).Error(errcode.ErrParams.WithCause(err).AppendMsg("参数错误"))
		return
	}
	if err = c.svc.Register(ctx, &user); err != nil {
		app.NewResponse(ctx).Error(err)
		return
	}
	app.NewResponse(ctx).SuccessOk()
}
