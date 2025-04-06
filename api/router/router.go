package router

import (
	"back/api/controller"
	"github.com/gin-gonic/gin"
)

// 路由注册相关的, 都放在router 包下

func RegisterRoutersAndMiddleware(q *controller.QuestionController, fs ...gin.HandlerFunc) *gin.Engine {
	s := gin.Default()
	RegisterMiddleware(s, fs...)

	g := s.Group("/api/v1")
	registerQuestion(g, q)
	return s
}

func RegisterMiddleware(g *gin.Engine, fs ...gin.HandlerFunc) *gin.Engine {
	g.Use(fs...)
	return g
}
