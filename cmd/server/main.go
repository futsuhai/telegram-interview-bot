package main

import (
	"log"

	"github.com/futsuhai/telegram-interview-bot/internal/config"
)

func main() {
	cfg, vaultClient, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Println("Starting Telegram bot server...")
	log.Println("Bot token loaded:", cfg.TelegramToken)

	if vaultClient != nil {
		log.Println("Vault client initialized")
	}
}
