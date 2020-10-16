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

	models.DB.Model(&models.Failed{}).Where("resend = ?", "0")
	//GET
	app.GET("/send",handlers.Send)
	app.GET("/panel",handlers.Panel)

	//POST
	app.POST("/search",handlers.SearchPhone)

	//Not found
	app.NoRoute(handlers.NotFound)

	app.Run(":3000")
}