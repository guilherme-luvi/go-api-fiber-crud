package config

import (
	"os"

	"github.com/guilherme-luvi/go-api-fiber-crud/pkg/database"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	Port           string
	StarWarsAPIURL string
	SecretKey      []byte
	logger         *Logger
	db             *gorm.DB
)

func InitEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	StarWarsAPIURL = os.Getenv("STAR_WARS_API_URL")
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	return nil
}

func InitDB() error {
	var err error

	db, err = database.InitSQLite()
	if err != nil {
		return err
	}

	logger.Info("Database connection established successfully")
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

func GetHttpClient() *HttpClientService {
	return NewHttpClientService()
}
