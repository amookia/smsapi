package models


import "gorm.io/gorm"

type Success struct {
	gorm.Model
	PhoneNumber  int
	SmsBody string
	Time int
	Api string
}

type Failed struct {
	gorm.Model
	Number int
	Body string
	Api string
	Resend bool
}