package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func extractData(msg string) (error, string, string, int, int, string) {

	data := strings.Split(msg, "|")
	what := data[0]
	what_type := detectType(what)
	price := calculatePrice(data[1])

	if price == 0 {
		return fmt.Errorf("Giá tiền bị sai roài!"), "", "", 0, 0, ""
	}

	quantity, err := strconv.Atoi(data[2])
	if err != nil {
		return err, "", "", 0, 0, ""
	}

	var timestamp string
	if len(data) == 4 {
		timestamp = data[3]
	} else {
		timestamp = time.Now().Format("02/01/2006")
	}
	return nil, what, what_type, price, quantity, timestamp
}
