package main

import (
	"log/slog"
	"os"

	"github.com/myevi/go-biwywfok/pkg/app"
	"github.com/myevi/go-biwywfok/pkg/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger()
	log = log.With(slog.String("env", envLocal))

	log.Info("biwywfok started")
	log.Debug("logger debug mode enabled")
	app.Start(cfg)
}

func setupLogger() *slog.Logger {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log
}
