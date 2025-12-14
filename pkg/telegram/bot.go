package telegram

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBot() (*tgbotapi.BotAPI, error) {
	token := os.Getenv("TELEGRAM_APITOKEN")
	if token == "" {
		return nil, fmt.Errorf("TELEGRAM_APITOKEN not set")
	}

	return tgbotapi.NewBotAPI(token)
}
