package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/example/go-starter/internal/config"
	"github.com/example/go-starter/internal/db"
	"github.com/example/go-starter/internal/logger"
	"github.com/example/go-starter/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Load()

	logg, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	logs := logg.Sugar()

	defer func() { _ = logg.Sync() }()

	sqlDB, err := db.New(cfg, logs)
	if err != nil {
		logs.Fatalf("db init: %v", err)
	}
	defer sqlDB.Close()

	app := server.New(cfg, logs, sqlDB)

	go func() {
		addr := ":" + cfg.AppPort
		logs.Infof("listening on %s", addr)
		if err := app.Listen(addr); err != nil {
			logs.Fatalf("server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logs.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := app.Shutdown(); err != nil {
		logs.Warnf("fiber shutdown error: %v", err)
	}
	<-shutdownCtx.Done()
	logs.Info("graceful shutdown complete")
}
