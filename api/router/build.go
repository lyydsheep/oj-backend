package router

func registerBuild(s *gin.RouterGroup, build *controller.BuildController) {
	g := s.Group("/build")
	g.GET("/pagination", build.TestPagination)
	g.GET("/test_gorm_log", build.TestGormLog)
	g.POST("/test_create", build.CreateDemoOrder)
}
