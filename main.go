package main

func init() {
	config.InitConfig()
	log.InitLogger()
	dao.InitGormLogger()
	dao.InitDB()
}

func main() {
	if config.App.Env == enum.ModePROD {
		gin.SetMode(gin.ReleaseMode)
	}
	g := InitializeApp()
	if err := g.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
