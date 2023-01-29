package server

import (
	"transaction/config"
	"transaction/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIService(r *gin.Engine){
	var db *gorm.DB
	db = config.GetDBConnection()

	transRepo := domain.NewTransactionRepository(db)
	transUsecase := domain.NewTransactionUsecase(transRepo)
	transController := domain.NewTransactionController(transUsecase)

	registerTransactionRoute(r, transController)
}