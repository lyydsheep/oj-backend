package router

import (
	"back/api/controller"
	"github.com/gin-gonic/gin"
)

func registerUser(s *gin.RouterGroup, u *controller.UserController) {
	g := s.Group("/user")
	g.POST("/login", u.Login)
	g.POST("/register", u.Register)
}
