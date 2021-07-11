package main

import (
	"fmt"
	"log"

	"github.com/ikaruswill/koffea/client/telegram"
	"github.com/ikaruswill/koffea/config"
	"github.com/ikaruswill/koffea/koffea"
	"github.com/ikaruswill/koffea/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	koffeaConfig := &config.Config{}
	err := koffeaConfig.Load("config.yaml")

	if err != nil {
		log.Fatal(err)
		return
	}

	var store storage.Storage
	store, err = storage.NewStorage(koffeaConfig.Storage)
	if err != nil {
		log.Fatal(err)
		return
	}

	tClient, err := telegram.NewConnection(koffeaConfig.Telegram)
	if err != nil {
		log.Fatal(err)
		return
	}

	app := koffea.Koffea{
		UserService:     storage.NewUserService(store),
		OrderService:    storage.NewOrderService(store),
		GroupBuyService: storage.NewGroupBuyService(store),
		TelegramService: tClient,
	}

	tClient.Bot.Handle("/start", func(m *tb.Message) {
		user := &koffea.User{
			FirstName:    m.Sender.FirstName,
			LastName:     m.Sender.LastName,
			Username:     m.Sender.Username,
			LanguageCode: m.Sender.LanguageCode,
			IsBot:        m.Sender.IsBot,
		}
		app.UserService.CreateIfNotExists(user)

		tClient.Bot.Send(m.Sender, fmt.Sprintf("Hi %s! My name is Koffea, how can I help you today?", user.FirstName))
		tClient.Bot.Send(m.Sender, `/listgb - List active group buys
		/viewgb - View information about a group buy
		/startgb - Start a group buy
		/joingb - Join a group buy
		/endgb - End a group buy`)
	})

	tClient.Bot.Start()
}
