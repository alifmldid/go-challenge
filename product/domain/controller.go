package domain

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController interface{
	Insert(c *gin.Context)
	FindById(c *gin.Context)
}

type productController struct{
	productUsecase ProductUsecase
}

func NewProductController(productUsecase ProductUsecase) ProductController {
	return &productController{productUsecase: productUsecase}
}

func (controller *productController) Insert(c *gin.Context) {
	var product Product
	c.ShouldBindJSON(&product)

	err := controller.productUsecase.Insert(c.Request.Context(), product)

	message := "success"
	if err != nil {
		message = "fail"
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}

func (controller *productController) FindById(c *gin.Context){
	message := "success"
	var data Product

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		message = "fail"
		panic(err.Error())
	}

	product, err :=  controller.productUsecase.FindById(c, id)

	if err != nil {
		message = "product not found"
		c.JSON(404, gin.H{
			"message": message,
		})
	} else {
		data = product
		c.JSON(200, gin.H{
			"message": message,
			"data": data,
		})
	}
}