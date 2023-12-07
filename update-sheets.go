package main

import (
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strconv"
	"time"
)

func updateSheets(user string, what string, what_type string, quantity int, price int, timestamp string) (string, error) {

	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	url := goDotEnvVariable("WEBHOOK_URL")
	//add url search params
	url += "&what=" + url2.QueryEscape(what)
	url += "&what_type=" + url2.QueryEscape(what_type)
	url += "&quantity=" + url2.QueryEscape(strconv.Itoa(quantity))
	url += "&price=" + url2.QueryEscape(strconv.Itoa(price))
	url += "&time=" + url2.QueryEscape(timestamp)
	url += "&time_detail=" + url2.QueryEscape(time.Now().In(loc).Format("02/01/2006 15:04:05"))
	url += "&user=" + url2.QueryEscape(user)
	Info.Println(url)
	response, err := http.Get(url)

	if err != nil {
		Error.Println(err)
		return "", err
	}

	response_data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Error.Println(err)
		return "", err
	}

	return string(response_data), nil
}
