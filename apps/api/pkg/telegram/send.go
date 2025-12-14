package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"log"
)

var telegramAPI = "https://api.telegram.org/bot%s/sendMessage"

type sendMessageRequest struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func sendMessage(chatID int64, text string) {
	token := os.Getenv("TELEGRAM_APITOKEN")
	if token == "" {
		panic("TELEGRAM_APITOKEN not set")
	}

	url := fmt.Sprintf(telegramAPI, token)

	payload := sendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}

	body, _ := json.Marshal(payload)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		// Log only — never panic in serverless
		log.Println("telegram send error: %w", err)
	}
}
