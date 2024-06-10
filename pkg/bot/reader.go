package bot

import (
	"errors"
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (tg *TelegramBot) reader(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	var responseMessage string

	switch msg.Text {
	case "/start":
		responseMessage = "hi dude"
		fmt.Println("\033[91mmessage\033[0m", slog.String("from", msg.From.UserName), slog.String("message", msg.Text))
	default:
		return nil, errors.New("unknown message text")
		//todo здесь добавить логику проверки на наличие контекста
	}
	response := tgbotapi.NewMessage(msg.Chat.ID, responseMessage)
	// response.ReplyToMessageID = msg.MessageID

	return &response, nil
}
