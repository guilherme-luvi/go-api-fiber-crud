package config

import (
	"os"

	"github.com/guilherme-luvi/go-api-fiber-crud/src/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "././db/main.db"

	// check if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file does not exist, creating...")

		// create db directory
		err = os.MkdirAll("././db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		// create db file
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
		logger.Info("database file created successfully")
	}

	// connect to sqlite database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to sqlite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		logger.Errorf("Error migrating schemas: %v", err)
		return nil, err
	}

	logger.Info("Database connection established successfully")
	return db, nil
}
