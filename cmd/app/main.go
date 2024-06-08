package main

import (
	"github.com/guilherme-luvi/go-api-fiber-crud/cmd/router"
	"github.com/guilherme-luvi/go-api-fiber-crud/pkg/config"
)

func main() {
	// init logger
	logger := config.GetLogger("main")

	// init env variables
	if err := config.InitEnvVariables(); err != nil {
		logger.Error("Error loading .env file")
		return
	}

	// init database
	err := config.InitDB()
	if err != nil {
		logger.Error("Error initializing database")
		return
	}

	router.SetupRouter()
}
