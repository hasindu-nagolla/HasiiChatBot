package main

import (
	"log"
	"time"
	"github.com/hasindu-nagolla/HasiiChatBot/config"
	"github.com/hasindu-nagolla/HasiiChatBot/database"
	"github.com/hasindu-nagolla/HasiiChatBot/handlers"
	tele "gopkg.in/telebot.v3"
)

func main() {
	// get env vars
	cfg := config.LoadConfig()

	// connect to mongo
	db := database.ConnectDB(cfg.MongoURL)
	defer db.Client.Disconnect(nil)

	// init telebot
	pref := tele.Settings{
		Token:  cfg.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// attach routes
	handlers.RegisterCommands(b, db)
	handlers.RegisterChatbot(b, db)

	// run
	log.Println("Bot started successfully in Golang (Clean Architecture)!")
	b.Start()
}
