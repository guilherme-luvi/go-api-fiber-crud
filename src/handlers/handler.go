package handlers

import (
	"github.com/guilherme-luvi/go-api-fiber-crud/src/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handlers")
	db = config.GetDB()
}
