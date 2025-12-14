package api

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ximura/giftweaver/pkg/repository"
	"github.com/ximura/giftweaver/pkg/telegram"
)

var (
	pool *repository.Repository
	once sync.Once
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	p, err := getPool(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Telegram webhook always sends POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusOK)
		return
	}

	bot, err := telegram.NewBot()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	queries := p.Queries
	if err := telegram.HandleWebhook(ctx, bot, queries, &update); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getPool(ctx context.Context) (*repository.Repository, error) {
	var err error

	once.Do(func() {
		pool, err = repository.NewRepository(ctx)
	})

	return pool, err
}
