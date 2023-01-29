package main

import (
	"auth/server"

	"github.com/gin-gonic/gin"
)

func main(){
	var r *gin.Engine
	r = gin.Default()

	server.RegisterAuthService(r)

	r.Run(":8080")
}