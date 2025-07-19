package main

import (
	"log"
	"os"

	"discord-bot-go/bot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_TOKEN_ID")
	if token == "" {
		log.Fatal("DISCORD_TOKEN_ID not set in .env")
	}

	bot.BotToken = token
	bot.Run()
}
