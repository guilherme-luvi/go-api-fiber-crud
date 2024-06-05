package main

import (
	"github.com/guilherme-luvi/go-api-fiber-crud/src/config"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/router"
)

func main() {
	// init logger
	logger := config.GetLogger("main")

	// init env variables
	if err := config.InitEnvVariables(); err != nil {
		logger.Error("Error loading .env file")
		return
	}

	// Setup router
	router.SetupRouter()
}