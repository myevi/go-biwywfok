package entities

import "time"

type TelegramUser struct {
	ID             int
	TelegramUserID uint64
	Username       string
	FirstName      string
	LastName       string
	LanguageCode   string
	CreatedAt      time.Time
}

type TelegramMessage struct {
	ID             int
	UserID         int
	TelegramChatID uint64
	Text           string
	CreatedAt      time.Time
}
