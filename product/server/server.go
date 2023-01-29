package server

import (
	"product/config"
	"product/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAPIService(r *gin.Engine){
	var db *gorm.DB
	db = config.GetDBConnection()

	productRepo := domain.NewProductRepository(db)
	productUsecase := domain.NewProductUsecase(productRepo)
	productController := domain.NewProductController(productUsecase)

	registerProductRoute(r, productController)
}