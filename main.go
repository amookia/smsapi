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

	//register routes
	handlers.RegisterFront(app)
	handlers.RegisterApi(app)

	//Not found
	app.NoRoute(handlers.NotFound)

	app.Run(":3000")
}