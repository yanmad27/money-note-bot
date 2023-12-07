package main

import (
	"fmt"
	"github.com/Pramod-Devireddy/go-exprtk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func validateCommand(msg string) bool {

	msg1 := msg
	regex1 := regexp.MustCompile(`^/note\s(.+)\|(.+)\|(\d+)\|(.+)$`)
	valid1 := regex1.MatchString(msg1)
	fmt.Println("valid1", msg1, valid1)

	msg2 := msg
	regex2 := regexp.MustCompile(`^/note\s(.+)\|(.+)\|(\d+)$`)
	valid2 := regex2.MatchString(msg2)
	fmt.Println("valid2", msg2, valid2)
	return valid1 || valid2
}

func calculatePrice(price string) int {
	exprtkObj := exprtk.NewExprtk()

	exprtkObj.SetExpression(price)

	err := exprtkObj.CompileExpression()
	if err != nil {
		return 0
	}

	exprtkObj.SetDoubleVariableValue("x", 8)
	return int(exprtkObj.GetEvaluatedValue())

}

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

func Format(n int) string {
	in := strconv.Itoa(n)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = '.'
		}
	}
}

func detectType(what string) string {
	what = strings.ToLower(what)

	food_keyword := []string{"bbq", "food", "rau", "nước", "củ", "trái", "bánh", "kem", "mì", "bún", "phở", "cơm", "cháo", "chè", "cá", "thịt"}
	for _, keyword := range food_keyword {
		if strings.Contains(what, keyword) {
			return "Food"
		}
	}

	medicine_keyword := [5]string{"thuốc", "viên", "hộp", "gói", "chai"}
	for _, keyword := range medicine_keyword {
		if strings.Contains(what, keyword) {
			return "Thuốc"
		}
	}

	return "Gia dụng"
}

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

func main() {

	bot, err := tgbotapi.NewBotAPI(goDotEnvVariable("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	Info.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(-1)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	init_message := "Hello my bae 🥰\nDùng cú pháp bên dưới nha\nCP: /note what|price|quantity|time\nVD1: /note bánh mì|10000*2+5000*3|2\nVD2: /note bánh mì|10000*2+5000*3|2|07/11/1998"

	for update := range updates {
		if update.Message != nil && !update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = init_message
			if _, err := bot.Send(msg); err != nil {
				Error.Println(err)
			}
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "note":
			if !validateCommand(update.Message.Text) {
				msg.Text = "Sai cú pháp rồi bae 🥺\nPhải như zị nè:\n/note bánh mì|10000*2+5000*3|2|07/11/1998"
				if _, err := bot.Send(msg); err != nil {
					Error.Println(err)
				}
				continue
			}

			err, what, what_type, price, quantity, timestamp := extractData(update.Message.CommandArguments())
			if err != nil {
				msg.Text = "Oops, em nhập sai gì đó rồi 🥺\nThử lại nha\n"
				msg.Text += err.Error()
				if _, err := bot.Send(msg); err != nil {
					Error.Println(err)
				}
				continue
			}

			total := price * quantity

			msg.Text = "Okie, a đang note đợi tí nha, hihi 😚"
			if _, err := bot.Send(msg); err != nil {
				Error.Println(err)
			}

			response, err := updateSheets(update.Message.From.UserName, what, what_type, quantity, price, timestamp)

			if err != nil {
				Error.Println(err)
				msg.Text = "Oops, có lỗi xảy ra rồi bae 🥺\nThử lại sau nhé"
			} else {
				Info.Println(response)
				msg.Text = "Xong rùi nè ❤️\n" + "Em mua " + Format(quantity) + " " + what + "(" + what_type + ")" + " giá " + Format(price) + " đồng.\nTổng là: " + Format(price) + " x " + Format(quantity) + " = " + Format(total) + " đồng"

			}
		default:
			msg.Text = init_message
		}

		if _, err := bot.Send(msg); err != nil {
			Error.Println(err)
		}
	}
}
