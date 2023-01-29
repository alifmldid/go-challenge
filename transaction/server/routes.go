package server

import (
	"fmt"
	"net/http"
	"strings"
	"transaction/domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func registerTransactionRoute(r *gin.Engine, transactionController domain.TransactionController){
	transaction := r.Group("/")
	transaction.Use(AuthRequired)
	{
		transaction.POST("/transaction", transactionController.Insert)
		transaction.GET("/transaction/:id", transactionController.FindById)
	}
}

func AuthRequired(c *gin.Context){
	authorizationHeader := c.Request.Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte("SECRET_NUMBER"), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Set("token", token.Raw)
	c.Next()
}