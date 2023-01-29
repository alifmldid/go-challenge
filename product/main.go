package main

import (
	"product/server"

	"github.com/gin-gonic/gin"
)

func main(){
	var r *gin.Engine
	r = gin.Default()
	
	server.RegisterAPIService(r)

	r.Run(":8000")
}