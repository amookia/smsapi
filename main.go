package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"smsproject/handlers"
	"smsproject/models"
)

func main(){
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")

	//DB
	models.ConnectToDB()
	xxx,err := models.TopNum(9031900415)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(xxx)
	//GET
	app.GET("/send",handlers.Send)
	app.GET("/panel",handlers.Panel)

	//POST
	app.POST("/search",handlers.SearchPhone)

	app.Run(":3000")
}