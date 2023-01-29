package server

import (
	"auth/config"
	"auth/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthService(r *gin.Engine){
	var db *gorm.DB

	db = config.GetDBConnection();

	authRepo := domain.NewAuthRepository(db)
	authUsecase := domain.NewAuthUsecase(authRepo)
	authController := domain.NewAuthController(authUsecase)

	registerAuthRoute(r, authController)
}