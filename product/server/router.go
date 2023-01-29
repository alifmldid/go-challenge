package server

import (
	"fmt"
	"net/http"
	"product/domain"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func registerProductRoute(r *gin.Engine, productController domain.ProductController){
	product := r.Group("/")
	product.Use(AuthRequired)
	{
		product.POST("/product", productController.Insert)
		product.GET("/product/:id", productController.FindById)
	}
}

func AuthRequired(c *gin.Context){
	authorizationHeader := c.Request.Header.Get("Authorization")
	fmt.Println(authorizationHeader)

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
	
	c.Next()		
}