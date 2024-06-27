package entities

type ChatGptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
