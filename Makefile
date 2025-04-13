# ===== CONFIG =====
BINARY_NAME=bot
MAIN_PACKAGE=./cmd/bot

# Export environment variables in the Makefile
export GEMINI_API_KEY=dummy
export TG_BOT_TOKEN=dummy
export WEBHOOK_URL=dummy
# ===== TARGETS =====

# Build the bot binary
build:
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Run the bot, environment vars are already set
run: build
	./$(BINARY_NAME)

# Clean the project
clean:
	rm -f $(BINARY_NAME)
