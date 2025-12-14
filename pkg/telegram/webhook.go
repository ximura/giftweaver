package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ximura/giftweaver/pkg/repository"
)

type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
	Chat      Chat   `json:"chat"`
	From      User   `json:"from"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
}

func HandleWebhook(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	q *repository.Queries,
	update *tgbotapi.Update,
) error {
	if update.Message == nil {
		return nil
	}

	msg, err := HandleMessage(update)
	if err != nil {
		return err
	}
	_, err = bot.Send(msg)
	return err
}
