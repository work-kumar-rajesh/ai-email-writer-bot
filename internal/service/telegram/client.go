package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// TelegramBotClient defines methods to interact with Telegram.
type TelegramBotClient interface {
	SendMessage(chatID int64, message string) ([]*tgbotapi.Message, error)
}
