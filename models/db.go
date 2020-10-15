package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)




var DB *gorm.DB

func ConnectToDB(){
	dsn := "root:kialoloo@tcp(127.0.0.1:3306)/smsapi?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("FAILED TO CONECT DATABASE!")
	}
	db.AutoMigrate(&Success{})
	db.AutoMigrate(&Failed{})

	fmt.Println("migraaaaaaaaaaaaaate")

	DB = db

}