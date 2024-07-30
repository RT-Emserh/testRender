package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(ctx *gin.Engine) {
	ctx.GET("/hello", Controller)
}

func main() {
	server := gin.Default()
	Router(server)
	server.Run()
}

func Controller(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"mensagge": "hello word",
	})
}
