package domain

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface{
	Login(c *gin.Context)
}

type authController struct{
	authUsecase AuthUsecase
}

func NewAuthController(authUsecase AuthUsecase) AuthController{
	return &authController{authUsecase}
}

func (controller *authController) Login(c *gin.Context){
	var authPayload AuthPayload
	c.ShouldBindJSON(&authPayload)
	code, message, err := controller.authUsecase.Login(c, authPayload)

	if err != nil {
		panic(err)
	}

	if code != 200 {
		c.JSON(code, gin.H{
			"message": message,
		})
	} else {
		c.JSON(code, gin.H{
			"token": message,
		})
	}
}