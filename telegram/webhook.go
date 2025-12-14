package telegram

import (
	"context"

	"github.com/ximura/giftweaver/internal/repository"
)

type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

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
	q *repository.Queries,
	update Update,
) error {
	if update.Message == nil {
		return nil
	}

	return HandleMessage(update.Message)
}
