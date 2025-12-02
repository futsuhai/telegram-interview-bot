package main

import (
	"log"

	"github.com/futsuhai/telegram-interview-bot/internal/config"
)

func main() {
	// 1. Загружаем конфигурацию + Vault клиент
	cfg, vaultClient, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Println("Starting Telegram bot server...")
	log.Println("Bot token loaded:", cfg.TelegramToken)

	// 2. Vault клиент может быть nil в dev, но готов к использованию
	if vaultClient != nil {
		log.Println("Vault client initialized")
		// Пример: можно читать дополнительные секреты
		// secret, err := vaultClient.Logical().Read("secret/data/other")
	}

	// 3. Создаём Telegram клиента и сервисы
	// Здесь будет запуск бота
}
