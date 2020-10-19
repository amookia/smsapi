package handlers


var SendForm struct{
	Phone int `form:"number" binding:"required"`
	Body string `form:"body" binding:"required"`
}

var SearchForm struct{
	Phone int `form:"phone" binding:"required"`
}