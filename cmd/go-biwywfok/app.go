package main

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/myevi/go-biwywfok/pkg/adapter/chatgpt"
	"github.com/myevi/go-biwywfok/pkg/bot"
)

type Config struct {
	TelegramToken string `yaml:"telegram_token"`
	ChatgptToken  string `yaml:"chatgpt_token"`
	ChatgptURL    string `yaml:"chatgpt_url"`
	ChatgptModel  string `yaml:"chatgpt_model"`
}

type App struct {
	bot *bot.TelegramBot
}

func start(cfg *Config) (err error) {
	openai, err := chatgpt.New(chatgpt.Config{
		Token: cfg.ChatgptToken,
		URL:   cfg.ChatgptURL,
		Model: cfg.ChatgptModel,
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

func mustLoad() *Config {
	configPath := "configs/config.yml"
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}
