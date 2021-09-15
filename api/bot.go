package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func StartSotaBot(tokenS string) (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	_ = godotenv.Load()
	token := os.Getenv(tokenS)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("error starting: %s", err.Error())
		return nil, nil
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Printf("error starting: %s", err.Error())
	}
	return bot, updates
}
