package app

import (
	"fmt"
	"log/slog"

	"github.com/myevi/go-biwywfok/pkg/bot"
	"github.com/myevi/go-biwywfok/pkg/config"
)

type App struct {
	bot *bot.TelegramBot
}

func Start(cfg *config.Config) {
	bot, err := bot.New(bot.Config{
		Token: cfg.TelegramToken,
	})

	if err != nil {
		fmt.Println("bot is not created", slog.String("error", err.Error()))
		return
	}

	app := &App{
		bot: bot,
	}

	app.bot.ReadMessages()
}
