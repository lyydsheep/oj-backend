//go:build wireinject

package main

import (
	"back/api/controller"
	"back/api/router"
	"back/common/middleware"
	"back/dal/dao"
	"back/logic/appService"
	"back/logic/domainService"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() *gin.Engine {
	wire.Build(router.RegisterRoutersAndMiddleware,
		middleware.GetHandlerFunc, controller.NewQuestionController,
		appService.NewQuestionAppServiceV1,
		domainService.NewQuestionDomainServiceV1,
		dao.DBMaster)
	return nil
}
