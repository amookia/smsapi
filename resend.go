package main

import (
	"fmt"
	"log"
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
			//body := sms["body"]
			number, _ := sms["number"].(string)
			id := sms["id"].(uint)
			updateval(fmt.Sprint(id), number)
		}
		time.Sleep(5 * time.Minute)
	}
}

func resend(number uint,body string,id int){
	url := "/send?number=" + fmt.Sprint(number) + "&body=" + url2.QueryEscape(body)
	api,err := http.Get("http://localhost:81" + url)
	if err != nil || api.StatusCode != 200{
		api2,err2 := http.Get("http://localhost:82" + url)
		if err2 != nil || api2.StatusCode != 200 {
			log.Println(err)
		}else{
			fmt.Println("OK")
		}
	}else{
		fmt.Println("OK")
	}
}

func updateval(id string,number string){
	query := "id = ? AND resend = ?"
	models.DB.Model(&models.Failed{}).Where(query,id,"0").Update("resend","1")

}