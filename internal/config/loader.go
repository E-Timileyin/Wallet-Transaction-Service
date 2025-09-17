package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() error {
	// Load .env first
	_ = godotenv.Load()

	// Viper setup for optional config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No config.yaml found, relying on .env or system env")
	}

	return nil
}
