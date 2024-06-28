package bot

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/myevi/go-biwywfok/pkg/adapter/chatgpt"
	"github.com/myevi/go-biwywfok/pkg/entities"
)

type Config struct {
	Token string
	URL   string
}

type TelegramBot struct {
	Bot          *tgbotapi.BotAPI
	Config       Config
	OpenAI       *chatgpt.Client
	UserMessages []entities.ChatgptMessage
}

func New(cfg Config, chatgptClient *chatgpt.Client) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		Bot:          bot,
		Config:       cfg,
		OpenAI:       chatgptClient,
		UserMessages: make([]entities.ChatgptMessage, 1),
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
			// TODO why do i need chech that response?
			response, err := tg.Bot.Send(*tg.reader(update.Message))
			if err != nil {
				slog.Error("bot: read message", slog.String("error", err.Error()))
			}

			slog.Info("bot: read message", slog.String("msg", response.Text))
		}
	}
}
