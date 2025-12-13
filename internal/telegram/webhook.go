package telegram

import (
	"encoding/json"
	"io"
	"net/http"
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

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	var update Update
	if err := json.Unmarshal(body, &update); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// Telegram sometimes sends updates without message (callbacks, etc.)
	if update.Message == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	HandleMessage(update.Message)

	// Telegram requires a 200 OK
	w.WriteHeader(http.StatusOK)
}
