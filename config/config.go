package config

import (
	"log"

	"github.com/spf13/viper"
)

// Load environment variables
func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

// Get environment variable with default value
func GetEnv(key, defaultValue string) string {
	if value, exists := viper.Get(key).(string); exists {
		return value
	}
	return defaultValue
}
