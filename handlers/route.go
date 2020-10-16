package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"smsproject/models"
	"smsproject/request"
)

//Method get
//parameters number,body
func Send(c *gin.Context){
	//Check form
	formvalid := c.ShouldBind(&SendForm) == nil
	if formvalid {
		//Call sendsms function
		_,err := request.SendSms(SendForm.Phone,SendForm.Body)

		//Check failed
		if err != nil {
			c.JSON(400,gin.H{
				"message" : "Failed to send sms!",
			})
		}else {
			//If not failed!
			c.JSON(200,gin.H{"message":"SMS sent!","sent_to":SendForm.Phone})
		}
	}else {
		c.JSON(400,gin.H{"message":"invalid form"})
	}
}


func Panel(c *gin.Context) {
	//ALL SUCCESS COUNT
	var success,succuess_api1,succuess_api2 int64
	models.DB.Model(&models.Success{}).Count(&success)
	models.DB.Model(&models.Success{}).Where("api = ?", "api1").Count(&succuess_api1)
	models.DB.Model(&models.Success{}).Where("api = ?", "api2").Count(&succuess_api2)

	//ALL FAILED COUNT
	var failed1,failed2,total_failed int64
	models.DB.Model(&models.Failed{}).Count(&total_failed)
	models.DB.Model(&models.Failed{}).Where("api = ? AND resend = ?", "api1","0").Count(&failed1)
	models.DB.Model(&models.Failed{}).Where("api = ? AND resend = ?", "api2","0").Count(&failed2)

	//fmt.Println(float64(failed1)/float64(total_failed) * 100)
	//fmt.Println(float64(failed2)/float64(total_failed) * 100)


	c.HTML(200,"index.html",gin.H{
		"total":success,
		"total_success1":succuess_api1,
		"total_success2":succuess_api2,
		"api1_percentage":int(float64(failed1)/float64(total_failed) * 100),
		"api2_percentage":int(float64(failed2)/float64(total_failed) * 100),
	})
}

func SearchPhone (c *gin.Context) {
	isvalid := c.ShouldBind(&SearchForm) == nil
	if isvalid {
		var total int64
		fmt.Println(total)
		models.DB.Model(&models.Success{}).Where("phone_number = ?",SearchForm.Phone).Count(&total)
		if total == 0 {
			c.HTML(404,"errors.html",gin.H{"message":"NOT FOUND"})
			return
		}
		var results []map[string]interface{}

		models.DB.Model(&models.Success{}).Where("phone_number = ?",SearchForm.Phone).Find(&results)
		c.HTML(200, "search.html", gin.H{
			"message": "OK",
			"total":   total,
			"Lists": results,
		})
		return
	}else{
		c.HTML(400,"errors.html",gin.H{"message":"Invalid form"})
		return
	}
}