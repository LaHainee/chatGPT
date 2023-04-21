package main

import (
	openaiGw "chatgpt_bot/internal/gateway/openai"
	"chatgpt_bot/internal/model"
	"chatgpt_bot/internal/telegram/updates"
	"context"
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getUpdateConfig() tgbotapi.UpdateConfig {
	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 5

	return cfg
}

func main() {
	fmt.Println("Starting bot")

	ctx := context.Background()

	bot, err := tgbotapi.NewBotAPI(model.TelegramBotToken)
	if err != nil {
		panic(err)
	}

	updateConfig := getUpdateConfig()

	httpClient := &http.Client{}

	openai := openaiGw.NewGateway(model.OpenAiToken, httpClient)

	updatesHandler := updates.NewHandler(bot, openai)
	updatesHandler.HandleMessages(ctx, updateConfig)
}
