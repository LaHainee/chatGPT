package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Gateway struct {
	key    string
	client client
}

func NewGateway(key string, client client) *Gateway {
	return &Gateway{
		key:    key,
		client: client,
	}
}

func (g *Gateway) Chat(ctx context.Context, message string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	b, err := json.Marshal(chatRequest{
		Model: defaultModel,
		Messages: []chatMessage{
			{
				Role:    defaultRole,
				Content: message,
			},
		},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.key))
	req = req.WithContext(ctx)

	resp, err := g.client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("invalid http status")
	}

	var r chatResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}

	if len(r.Choices) == 0 {
		return "", errors.New("empty response")
	}

	return r.Choices[0].Message.Content, nil
}
