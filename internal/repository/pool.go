package repository

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrMissingDatabaseURL = errors.New("DATABASE_URL not set")

func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, ErrMissingDatabaseURL
	}

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	// 🔑 CRITICAL for serverless
	cfg.MaxConns = envInt32("PG_MAX_CONNS", 2)
	cfg.MinConns = envInt32("PG_IDLE_CONNS", 0)

	cfg.MaxConnLifetime = envDuration("PG_CONN_LIFETIME", 5*time.Minute)
	cfg.MaxConnIdleTime = 2 * time.Minute
	cfg.HealthCheckPeriod = 30 * time.Second

	// Fail fast on cold starts
	cfg.ConnConfig.ConnectTimeout = 5 * time.Second

	return pgxpool.NewWithConfig(ctx, cfg)
}

func envInt32(key string, def int32) int32 {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return int32(n)
		}
	}
	return def
}

func envDuration(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}
