package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(update *tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	text := strings.TrimSpace(update.Message.Text)

	switch {
	case text == "/start":
		return handleStart(update)

	case strings.HasPrefix(text, "/wish"):
		return handleWish(update.Message)

	default:
		return tgbotapi.NewMessage(update.Message.Chat.ID, "🎁 Send `/wish something` to add your gift wish."), nil
	}
}

func handleStart(update *tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	// Persist user + chat_id here
	// saveUser(msg.From.ID, msg.Chat.ID)

	// keyboard := tgbotapi.NewReplyKeyboard(
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("🛠 Створити кімнату"),
	// 		tgbotapi.NewKeyboardButton("🚪 Приєднатися до кімнати"),
	// 	),
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("ℹ️ Про гру"),
	// 	),
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("⚙️ Управління кімнатою"),
	// 	),
	// )

	keyboard2 := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("❌ Видалити", "remove_user:123"),
			tgbotapi.NewInlineKeyboardButtonData("↩️ Назад", "back"),
		),
	)

	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Обери потрібний варіант")
	msg.ReplyMarkup = keyboard2
	return msg, nil
}

func handleWish(msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	wish := strings.TrimSpace(strings.TrimPrefix(msg.Text, "/wish"))
	if wish == "" {
		return tgbotapi.NewMessage(msg.Chat.ID, "Please send `/wish <your gift wish>` 🎁"), nil
	}

	// saveWish(msg.From.ID, wish)
	return tgbotapi.NewMessage(msg.Chat.ID, "✅ Your wish has been saved!"), nil
}
