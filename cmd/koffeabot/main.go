package main

import (
	"log"
	"os"
	"time"

	"github.com/ikaruswill/koffeabot/internal/storage"
	"github.com/ikaruswill/koffeabot/internal/storage/user"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	dbClient := storage.Client{}
	err = dbClient.Init()

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hi! My name is Koffea, how can I help you today?")
		b.Send(m.Sender, `/listgb - List active group buys
		/viewgb - View information about a group buy
		/startgb - Start a group buy
		/joingb - Join a group buy
		/endgb - End a group buy`)

		dbClient.DB.Create(&user.User{
			Handle: m.Sender.Username,
		})
	})

	b.Start()
}
