package main

import (
	"log"

	"github.com/ikaruswill/koffea/internal/client/sqlite"
	"github.com/ikaruswill/koffea/internal/client/telegram"
	"github.com/ikaruswill/koffea/internal/config"
	"github.com/ikaruswill/koffea/internal/storage/user"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	koffeaConfig := &config.Config{}

	err := koffeaConfig.Load("config.yaml")

	if err != nil {
		log.Fatal(err)
		return
	}

	tClient, err := telegram.NewConnection(koffeaConfig.Telegram)

	if err != nil {
		log.Fatal(err)
		return
	}

	// var dbClient interface{}

	// switch koffeaConfig.Storage.Driver {
	// case config.SqliteStorageDriver:
	// 	dbClient, err = sqlite.NewConnection(koffeaConfig.Storage.Sqlite)
	// case config.PostgresStorageDriver:
	// 	dbClient, err = sqlite.NewConnection(koffeaConfig.Storage.Postgres)
	// }

	var dbClient *sqlite.Client
	dbClient, err = sqlite.NewConnection(koffeaConfig.Storage.Sqlite)

	if err != nil {
		log.Fatal(err)
		return
	}

	tClient.Bot.Handle("/start", func(m *tb.Message) {
		tClient.Bot.Send(m.Sender, "Hi! My name is Koffea, how can I help you today?")
		tClient.Bot.Send(m.Sender, `/listgb - List active group buys
		/viewgb - View information about a group buy
		/startgb - Start a group buy
		/joingb - Join a group buy
		/endgb - End a group buy`)

		dbClient.DB.Create(&user.User{
			FirstName:    m.Sender.FirstName,
			LastName:     m.Sender.LastName,
			Username:     m.Sender.Username,
			LanguageCode: m.Sender.LanguageCode,
			IsBot:        m.Sender.IsBot,
		})
	})

	tClient.Bot.Start()
}
