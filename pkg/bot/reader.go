package bot

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/myevi/go-biwywfok/pkg/entities"
)

var systemMessage = entities.ChatGptMessage{
	Role:    "system",
	Content: "You are helpful assistant",
}

func (tg *TelegramBot) reader(msg *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	slog.Info("\033[91m->message\033[0m", slog.String("from", msg.From.UserName), slog.String("text", msg.Text))
	var responseMessage string
	switch msg.Text {
	case "/start":
		tg.UserMessages = []entities.ChatGptMessage{
			systemMessage,
		}
		responseMessage = "ask me something"
	default:
		userMessage := entities.ChatGptMessage{
			Role:    "user",
			Content: msg.Text,
		}
		tg.UserMessages = append(tg.UserMessages, userMessage)
		// tg.OpenAI.ChatRequest(context.TODO(), tg.UserMessages)
		responseMessage = "some resp message uga buga"
	}
	response := tgbotapi.NewMessage(msg.Chat.ID, responseMessage)
	response.ReplyToMessageID = msg.MessageID

	slog.Info("\033[91mmessage->\033[0m", slog.String("to", msg.From.UserName), slog.String("text", responseMessage))
	return &response, nil
}
