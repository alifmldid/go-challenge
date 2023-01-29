package domain

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController interface{
	Insert(c *gin.Context)
	FindById(c *gin.Context)
}

type transactionController struct{
	transactionUsecase TransactionUsecase
}

func NewTransactionController(transUsecase TransactionUsecase) TransactionController {
	return &transactionController{transactionUsecase: transUsecase}
}

func (controller *transactionController) Insert(c *gin.Context) {
	var client = &http.Client{}
	var transaction Transaction
	c.ShouldBindJSON(&transaction)

	fmt.Println(transaction)
	prodId := strconv.Itoa(transaction.Product_id)

	request, err := http.NewRequest("GET", "http://localhost:8000/product/"+prodId, nil)
	if err != nil {
		panic(err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.JSON(404, gin.H{
			"message": "product not found",
		})
		return
	}
	
	err = controller.transactionUsecase.Insert(c.Request.Context(), transaction)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "fail",
		})
	}else{
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func (controller *transactionController) FindById(c *gin.Context){
	id := c.Param("id")
	transId, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	res, err := controller.transactionUsecase.FindById(c, transId)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		c.JSON(400, gin.H{
			"message": "fail",
		})
	}else{
		c.JSON(200, gin.H{
			"message": "success",
			"data": res,
		})
	}
}