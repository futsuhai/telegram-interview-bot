package config

import (
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string
}

func Load() (*Config, *vault.Client, error) {
	v := viper.New()

	vaultClient, err := initVault()
	if err != nil {
		return nil, nil, err
	}

	if err := loadFromVault(v, vaultClient); err != nil {
		return nil, vaultClient, err
	}

	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, vaultClient, err
	}

	return &cfg, vaultClient, nil
}

func initVault() (*vault.Client, error) {
	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultToken := os.Getenv("VAULT_TOKEN")

	if vaultAddr == "" || vaultToken == "" {
		log.Println("VAULT_ADDR or VAULT_TOKEN not set")
		return nil, nil
	}

	client, err := vault.NewClient(&vault.Config{Address: vaultAddr})
	if err != nil {
		return nil, err
	}
	client.SetToken(vaultToken)
	return client, nil
}

func loadFromVault(v *viper.Viper, client *vault.Client) error {
	secret, err := client.Logical().Read("secret/data/telegram-interview-bot")
	if err != nil {
		return err
	}

	if secret != nil {
		if data, ok := secret.Data["data"].(map[string]interface{}); ok {
			for key, value := range data {
				v.Set(key, value)
			}
		}
	}

	return nil
}
