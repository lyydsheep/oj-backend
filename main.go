package main

import (
	"back/common/enum"
	log "back/common/logger"
	"back/config"
	"back/dal/dao"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
	if err := g.Run(fmt.Sprintf(":%s", config.App.Port)); err != nil {
		panic(err)
	}
}
