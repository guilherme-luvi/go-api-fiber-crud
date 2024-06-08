package database

import (
	"log"
	"os"

	"github.com/guilherme-luvi/go-api-fiber-crud/internal/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"

	// check if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		log.Println("Database file does not exist")

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
		log.Println("Database file created successfully")
	}

	// connect to sqlite database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to sqlite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		log.Printf("Error migrating schemas: %v", err)
		return nil, err
	}

	log.Println("Database connection established successfully")
	return db, nil
}
