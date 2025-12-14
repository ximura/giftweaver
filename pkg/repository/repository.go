package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool    *pgxpool.Pool
	Queries *Queries
}

func NewRepository(ctx context.Context) (*Repository, error) {
	pool, err := NewPool(ctx)
	if err != nil {
		return nil, err
	}

	return &Repository{
		pool:    pool,
		Queries: New(pool),
	}, nil
}

func (p *Repository) Close() {
	// DO NOT close on Vercel
}
