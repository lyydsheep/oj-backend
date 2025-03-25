//go:build wireinject

package main

func InitializeApp() *gin.Engine {
	wire.Build(router.RegisterRoutersAndMiddleware,
		middleware.GetHandlerFunc, controller.NewBuildController,
		wire.Bind(new(appService.DemoAppService), new(*appService.DemoAppServiceV1)), appService.NewDemoAppServiceV1,
		wire.Bind(new(domainService.DemoDomainService), new(*domainService.DemoDomainServiceV1)), domainService.NewDemoDomainServiceV1,
		wire.Bind(new(dao.DemoDAO), new(*dao.DemoDAOV1)),
		dao.NewDemoDAO,
	)
	return nil
}
