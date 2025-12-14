package telegram

import (
	"strings"
)

func HandleMessage(msg *Message) error {
	text := strings.TrimSpace(msg.Text)

	switch {
	case text == "/start":
		return handleStart(msg)

	case strings.HasPrefix(text, "/wish"):
		return handleWish(msg)

	default:
		sendMessage(msg.Chat.ID, "🎁 Send `/wish something` to add your gift wish.")
	}

	return nil
}

func handleStart(msg *Message) error {
	// Persist user + chat_id here
	// saveUser(msg.From.ID, msg.Chat.ID)

	sendMessage(
		msg.Chat.ID,
		"🎅 Welcome to GiftWeaver!\n\nSend `/wish something` to add your gift wish.",
	)

	return nil
}

func handleWish(msg *Message) error {
	wish := strings.TrimSpace(strings.TrimPrefix(msg.Text, "/wish"))
	if wish == "" {
		sendMessage(msg.Chat.ID, "Please send `/wish <your gift wish>` 🎁")
		return nil
	}

	// saveWish(msg.From.ID, wish)

	sendMessage(msg.Chat.ID, "✅ Your wish has been saved!")
	return nil
}
