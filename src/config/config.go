package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Port   string
	logger *Logger
)

func InitEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
