package telegram

import (
	"strings"
)

func HandleMessage(msg *Message) {
	text := strings.TrimSpace(msg.Text)

	switch {
	case text == "/start":
		handleStart(msg)

	case strings.HasPrefix(text, "/wish"):
		handleWish(msg)

	default:
		sendMessage(msg.Chat.ID, "🎁 Send `/wish something` to add your gift wish.")
	}
}

func handleStart(msg *Message) {
	// Persist user + chat_id here
	// saveUser(msg.From.ID, msg.Chat.ID)

	sendMessage(
		msg.Chat.ID,
		"🎅 Welcome to GiftWeaver!\n\nSend `/wish something` to add your gift wish.",
	)
}

func handleWish(msg *Message) {
	wish := strings.TrimSpace(strings.TrimPrefix(msg.Text, "/wish"))
	if wish == "" {
		sendMessage(msg.Chat.ID, "Please send `/wish <your gift wish>` 🎁")
		return
	}

	// saveWish(msg.From.ID, wish)

	sendMessage(msg.Chat.ID, "✅ Your wish has been saved!")
}
