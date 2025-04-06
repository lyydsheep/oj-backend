package router

import (
	"back/api/controller"
	"github.com/gin-gonic/gin"
)

func registerQuestion(s *gin.RouterGroup, q *controller.QuestionController) {
	g := s.Group("/question")
	g.POST("", q.CreateQuestion)
}
