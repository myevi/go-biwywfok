package entities

import "time"

type Message struct {
	ID             int
	UserID         int
	TelegramChatID uint64
	Text           string
	CreatedAt      time.Time
}
