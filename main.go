package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatalf("%s", "no bot token provided")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("%s", err)
	}

	bot.Debug = true
	log.Println("Authorization username: ", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updatesChn := bot.GetUpdatesChan(updateConfig)

	for update := range updatesChn {
		if update.Message == nil {
			continue
		}

		log.Printf("new message: [%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf("Hello, %s", update.Message.From.UserName),
		)
		_, err := bot.Send(msg)
		if err != nil {
			log.Fatalf("error: %s", err.Error())
		}
	}
}
