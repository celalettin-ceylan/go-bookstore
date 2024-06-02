package utils

import (
	"log"

	"github.com/spf13/viper"
)

func ReadEnvFile() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
}
