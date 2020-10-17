package handlers

import "github.com/gin-gonic/gin"

func RegisterFront(app *gin.Engine){
	app.GET("/panel",Panel)
	app.POST("/search",SearchPhone)
}
func RegisterApi(app *gin.Engine){
	app.GET("/send",Send)
}
