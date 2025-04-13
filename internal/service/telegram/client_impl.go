package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// TelegramService contains the bot API and interacts with Telegram.
type TelegramService struct {
	bot *tgbotapi.BotAPI
}

// NewTelegramService initializes the Telegram bot service.
func NewTelegramService(botAPIKey string) TelegramService {
	bot, err := tgbotapi.NewBotAPI(botAPIKey)
	if err != nil {
		log.Fatalf("Error initializing bot: %s", err)
	}
	return TelegramService{
		bot: bot,
	}
}

// SetWebhook sets up the webhook URL for the bot.
func (t *TelegramService) SetWebhook(webhookURL string) (tgbotapi.APIResponse, error) {
	// Set the webhook for the bot
	webhook := tgbotapi.NewWebhook(webhookURL)
	return t.bot.SetWebhook(webhook)
}

// SendMessage sends a message to a specific chat on Telegram.
func (t TelegramService) SendMessage(chatID int64, message string) ([]*tgbotapi.Message, error) {
	const maxMessageLength = 4096
	var sentMessages []*tgbotapi.Message

	for start := 0; start < len(message); start += maxMessageLength {
		end := start + maxMessageLength
		if end > len(message) {
			end = len(message)
		}

		chunk := message[start:end]
		msg := tgbotapi.NewMessage(chatID, chunk)

		sentMessage, err := t.bot.Send(msg)
		if err != nil {
			return nil, err
		}
		sentMessages = append(sentMessages, &sentMessage)
	}

	return sentMessages, nil
}
