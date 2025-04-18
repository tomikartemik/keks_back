package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"tg_shop/internal/handler"
)

func AdmBotProcess(handlers *handler.AdminHandler) {
	botToken := os.Getenv("ADM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Admin bot token not provided")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to create admin bot: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("Received message: %s from user %d", update.Message.Text, update.Message.From.ID)
			switch update.Message.Text {
			case "/start":
				handlers.HandleAdminStart(bot, update)
			default:
				handlers.HandleAdminInput(bot, update)
			}
		} else if update.CallbackQuery != nil {
			log.Printf("Received callback query: %s from user %d", update.CallbackQuery.Data, update.CallbackQuery.From.ID)
			handlers.HandleCallbackQuery(bot, update.CallbackQuery)
		}
	}
}
