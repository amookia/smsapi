package request

import (
	"errors"
	"net/http"
	url2 "net/url"
	"smsproject/models"
	"strconv"
	"time"
)



//Checked [OK]
func SendSms (number int , body string) (status string,err error) {
		url := "/send?number=" + strconv.Itoa(number) + "&body=" + url2.QueryEscape(body)
		api,err := http.Get("http://localhost:81" + url)
		if err != nil {
			//
			failed(number,body,"api1")
			api2,err2 := http.Get("http://localhost:82" + url)

			if err2 != nil {
				//
				failed(number,body,"api2")
				return "",err2
			}
			if api2.StatusCode != 200 {
				return "",errors.New("sth is wrong!")
			}else {
				success(number,body,"api2")
				return "",nil
			}
		}else {
			if api.StatusCode != 200 {
				return "",errors.New("sth is wrong!")
			}else {
				success(number,body,"api1")
				return "",nil

			}
		}
}


//if status failed insert data into the fails
func failed(number int , body string , api string) {
	failedsms := models.Failed{
		Number: number,
		Body: body,
		Api : api,
		Resend: false,
	}
	models.DB.Create(&failedsms)
}

//if status ok insert data into the sms
func success(number int , body string , api string) {
	smstodb := models.Success{PhoneNumber: number ,
		SmsBody: body , Time: int(time.Now().Unix()) , Api: api}
	models.DB.Create(&smstodb)
}