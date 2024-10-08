package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
)

func main() {
	cfg := mustLoad()

	log := setupLogger()
	log = log.With(slog.String("env", envLocal))

	log.Info("biwywfok started")
	log.Debug("logger debug mode enabled")

	go func() {
		if err := start(cfg); err != nil {
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
