package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/myevi/go-biwywfok/pkg/entities"
)

type Request struct {
	Model    string                    `json:"model"` // TODO set constraints
	Messages []entities.ChatGptMessage `json:"messages"`
}

func (c *Client) ChatRequest(ctx context.Context, messages []entities.ChatGptMessage) (interface{}, error) {
	requestBody := Request{
		Model:    "gpt-3.5-turbo", // TODO make not a constant
		Messages: messages,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		c.baseURL,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	slog.Info("Request->", "body", string(body))
	response, err := c.http.Do(request)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read data from response body: %w", err)
	}

	slog.Info("response from openai", "data", string(data))
	os.Exit(1)
	if err != nil {
		return nil, fmt.Errorf("")
	}

	return struct{}{}, nil
}
