package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	aiagent "github.com/work-kumar-rajesh/ai-email-writer-bot/internal/service/ai-agent"
	"github.com/work-kumar-rajesh/ai-email-writer-bot/internal/service/telegram"
)

// TelegramHandler handles the incoming updates from Telegram.
type TelegramHandler struct {
	telegramService telegram.TelegramBotClient
	geminiService   aiagent.AIAgent
}

// NewTelegramHandler creates a new instance of TelegramHandler.
func NewTelegramHandler(telegramService telegram.TelegramBotClient, geminiService aiagent.AIAgent) *TelegramHandler {
	return &TelegramHandler{
		telegramService: telegramService,
		geminiService:   geminiService,
	}
}

// HandleUpdates processes the incoming updates from the webhook.
func (t *TelegramHandler) HandleUpdates(c *gin.Context) {
	var update tgbotapi.Update
	if err := c.ShouldBindJSON(&update); err != nil {
		log.Println("Failed to bind update:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if update.Message != nil {
		// Handle incoming message
		message := update.Message.Text
		chatID := update.Message.Chat.ID

		// Call Gemini API to generate a response (example for an email)
		response, err := t.geminiService.GenerateEmailReply(message)
		if err != nil {
			log.Println("Error generating email:", err)
			return
		}

		// Send the generated message back to the user
		_, err = t.telegramService.SendMessage(chatID, response)
		if err != nil {
			log.Println("Failed to send message:", err)
		}
	}
}
