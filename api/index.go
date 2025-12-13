package api

import (
	"net/http"

	"github.com/ximura/giftweaver/internal/telegram"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Process webhook payload (JSON from Telegram)
	telegram.HandleWebhook(w, r) // implement webhook logic
}
