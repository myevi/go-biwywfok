package bot

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	Token string
}

type TelegramBot struct {
	Bot    *tgbotapi.BotAPI
	Config Config
	//Must be DB here
}

func New(cfg Config) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		Bot:    bot,
		Config: cfg,
	}, nil
}

func (tg *TelegramBot) ReadMessages() {
	tg.Bot.Debug = false

	slog.Info("authorized on account", slog.String("name", tg.Bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			response, err := tg.reader(update.Message)
			if err != nil {
				slog.Error("bot: read message", slog.String("error", err.Error()))
				continue
			}

			tg.Bot.Send(*response)
		}
	}
}
