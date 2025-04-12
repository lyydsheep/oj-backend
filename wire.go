//go:build wireinject

package main

import (
	"back/api/controller"
	"back/api/router"
	"back/common/middleware"
	"back/dal/dao"
	"back/logic/repository"
	"back/logic/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() *gin.Engine {
	wire.Build(router.RegisterRoutersAndMiddleware,
		middleware.GetHandlerFunc, controller.NewQuestionController, controller.NewUserController,
		service.NewQuestionServiceV1, service.NewUserServiceV1,
		repository.NewQuestionRepositoryV1, repository.NewUserRepositoryV1,
		dao.DBMaster)
	return nil
}
