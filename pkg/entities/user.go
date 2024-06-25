package entities

import "time"

type User struct {
	ID             int
	TelegramUserID uint64
	Username       string
	FirstName      string
	LastName       string
	LanguageCode   string
	CreatedAt      time.Time
}