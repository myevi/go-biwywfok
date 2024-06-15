package bot

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (tg *TelegramBot) reader(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	slog.Info("\033[91m->message\033[0m", slog.String("from", msg.From.UserName), slog.String("text", msg.Text))
	var responseMessage string

	switch msg.Text {
	case "/start":
		responseMessage = "hi dude"
	default:
		responseMessage = "i dont understand"
		//todo здесь добавить логику проверки на наличие контекста
	}
	response := tgbotapi.NewMessage(msg.Chat.ID, responseMessage)
	response.ReplyToMessageID = msg.MessageID

	slog.Info("\033[91mmessage->\033[0m", slog.String("to", msg.From.UserName), slog.String("text", responseMessage))
	return &response, nil
}
