package updates

import (
	"context"
	"fmt"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot    bot
	openai openai
}

func NewHandler(b bot, o openai) *Handler {
	return &Handler{
		bot:    b,
		openai: o,
	}
}

func (h *Handler) HandleMessages(ctx context.Context, updateCfg tgbotapi.UpdateConfig) {
	updates := h.bot.GetUpdatesChan(updateCfg)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(updates tgbotapi.UpdatesChannel) {
		defer wg.Done()

		for update := range updates {
			if update.Message == nil {
				continue
			}

			fmt.Printf("Handling message from user: %d\n", update.Message.From.ID)

			// Получаем ответ от chatGPT на сообщение от пользователя
			respMessage, err := h.openai.Chat(ctx, update.Message.Text)
			if err != nil {
				_, _ = h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка, попроуйте еще раз"))
			}

			_, _ = h.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, respMessage))
		}
	}(updates)

	wg.Wait()
}
