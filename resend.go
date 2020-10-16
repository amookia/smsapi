package main

import (
	"fmt"
	"net/http"
	url2 "net/url"
	"smsproject/models"
	"time"
)

//Resend failed sms again
//sleep time 5 min
func main(){
	models.ConnectToDB()
	var notsent []map[string]interface{}
	for {
		models.DB.Model(&models.Failed{}).Where("resend = ?", "0").Find(&notsent)
		for _, sms := range notsent {
			body := sms["body"].(string)
			number, _ := sms["number"].(string)
			id := sms["id"].(uint)
			resend(fmt.Sprint(id), number,body)
		}
		time.Sleep(5 * time.Minute)
	}
}

func resend(id string,number string,body string){
	url := "/send?number=" + fmt.Sprint(number) + "&body=" + url2.QueryEscape(body)
	_,err := http.Get("http://localhost:81" + url)
	if err != nil{
		_,err2 := http.Get("http://localhost:82" + url)
		if err2 != nil {
			fmt.Println(err)
		}else{
			updateval(id, number)
		}
	}else{
		updateval(id, number)
	}
}

func updateval(id string,number string){
	query := "id = ? AND resend = ?"
	models.DB.Model(&models.Failed{}).Where(query,id,"0").Update("resend","1")

}