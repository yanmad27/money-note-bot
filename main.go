package main

import (
	//"github.com/Pramod-Devireddy/go-exprtk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
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
