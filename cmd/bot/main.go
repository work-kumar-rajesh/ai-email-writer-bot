package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/work-kumar-rajesh/ai-email-writer-bot/internal/handlers"
	"github.com/work-kumar-rajesh/ai-email-writer-bot/internal/service/ai-agent/gemini"
	"github.com/work-kumar-rajesh/ai-email-writer-bot/internal/service/telegram"
)

func main() {
	// Load environment variables from Makefile
	tgBotToken := os.Getenv("TG_BOT_TOKEN")
	geminiAPIKey := os.Getenv("GEMINI_API_KEY")
	webhookURL := os.Getenv("WEBHOOK_URL")

	// Validate env vars
	if tgBotToken == "" || geminiAPIKey == "" || webhookURL == "" {
		log.Fatal("Missing required environment variables: TG_BOT_TOKEN, GEMINI_API_KEY, or WEBHOOK_URL")
	}

	// Initialize Telegram bot service
	telegramService := telegram.NewTelegramService(tgBotToken)

	// Initialize Gemini service
	geminiService := gemini.NewGeminiService(geminiAPIKey)

	// Set webhook for Telegram
	_, err := telegramService.SetWebhook(webhookURL)
	if err != nil {
		log.Fatal("Failed to set webhook:", err)
	}

	// Create Telegram handler
	telegramHandler := handlers.NewTelegramHandler(telegramService, geminiService)

	// Set up Gin router
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		telegramHandler.HandleUpdates(c)
	})

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
