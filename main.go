package main

import (
	"github.com/gin-gonic/gin"
	"rabbitmq.com/gin/send"
)

type Email struct {
	Email string `json:"email"`
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}

var email Email

func Controller(ctx *gin.Context) {
	email.Email = "gerenciart.emserhml@gmail.com"
	d := email.Email
	send.Send(d)
	ctx.JSON(200, gin.H{
		"message": "Email sent",
	})
}

func main() {
	router := gin.Default()
	router.GET("/mensageria", Controller)
	router.Run(":8080") // specify the port you want to run the server on
}
