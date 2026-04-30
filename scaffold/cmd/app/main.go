package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yourname/go-starter/scaffold/internal/config"
	"github.com/yourname/go-starter/scaffold/internal/db"
	"github.com/yourname/go-starter/scaffold/internal/logger"
	"github.com/yourname/go-starter/scaffold/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Load()

	logg, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer func() { _ = logg.Sync() }()

	sqlDB, err := db.New(cfg, logg)
	if err != nil {
		logg.Sugar().Fatalf("db init: %v", err)
	}
	defer sqlDB.Close()

	app := server.New(cfg, logg, sqlDB)

	go func() {
		addr := ":" + cfg.AppPort
		logg.Sugar().Infof("listening on %s", addr)
		if err := app.Listen(addr); err != nil {
			logg.Sugar().Fatalf("server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logg.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := app.Shutdown(); err != nil {
		logg.Sugar().Warnf("fiber shutdown error: %v", err)
	}
	<-shutdownCtx.Done()
	logg.Info("graceful shutdown complete")
}
