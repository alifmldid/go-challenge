package server

import (
	"auth/domain"

	"github.com/gin-gonic/gin"
)

func registerAuthRoute(r *gin.Engine, authController domain.AuthController){
	auth := r.Group("/user")
	auth.POST("/login", authController.Login)
}