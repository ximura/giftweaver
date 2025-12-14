package webhook

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/ximura/giftweaver/internal/repository"
	"github.com/ximura/giftweaver/telegram"
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

	var update telegram.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	queries := p.Queries

	if err := telegram.HandleWebhook(ctx, queries, update); err != nil {
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
