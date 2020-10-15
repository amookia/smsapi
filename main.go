package main

import (
	"github.com/gin-gonic/gin"
	"smsproject/handlers"
	"smsproject/models"
)

func main(){
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")

	//DB
	models.ConnectToDB()

	//GET
	app.GET("/send",handlers.Send)
	app.GET("/panel",handlers.Panel)

	//POST
	app.POST("/search",handlers.SearchPhone)

	app.Run(":3000")
}