package db

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/yourname/go-starter/scaffold/internal/config"
)

// New creates a sqlx DB using pgx driver and performs a health ping
func New(cfg config.Config, logger *zap.Logger) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(context.Background(), "pgx", cfg.DatabaseURL)
	if err != nil {
		logger.Sugar().Errorf("connect db: %v", err)
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
