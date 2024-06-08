package handlers

import (
	"github.com/guilherme-luvi/go-api-fiber-crud/pkg/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandlers() {
	logger = config.GetLogger("handlers")
	db = config.GetDB()
}
