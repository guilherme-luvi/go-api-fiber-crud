package handlers

import "github.com/guilherme-luvi/go-api-fiber-crud/src/config"

var (
	logger *config.Logger
)

func InitHandler() {
	logger = config.GetLogger("handlers")
}
