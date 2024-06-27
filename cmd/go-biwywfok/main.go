package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/myevi/go-biwywfok/pkg/app"
	"github.com/myevi/go-biwywfok/pkg/config"
)

const (
	envLocal = "local"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger()
	log = log.With(slog.String("env", envLocal))

	log.Info("biwywfok started")
	log.Debug("logger debug mode enabled")

	go func() {
		if err := app.Start(cfg); err != nil {
			slog.Error("start app", "error", err)
			return
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	slog.Info(fmt.Sprint(<-ch))
	slog.Info("Stopping API server.")
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
