package app

import (
	"github.com/myevi/go-biwywfok/pkg/adapter/openai"
	"github.com/myevi/go-biwywfok/pkg/bot"
	"github.com/myevi/go-biwywfok/pkg/config"
)

type App struct {
	bot *bot.TelegramBot
}

func Start(cfg *config.Config) (err error) {
	openai, err := openai.New(openai.Config{
		Token: cfg.OpenAIToken,
		URL:   cfg.OpenAIURL,
	})
	if err != nil {
		return err
	}

	bot, err := bot.New(bot.Config{
		Token: cfg.TelegramToken,
	}, openai)
	if err != nil {
		return err
	}

	app := &App{
		bot: bot,
	}

	app.bot.ReadMessages()
	return
}
