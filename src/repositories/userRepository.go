package repositories

import (
	"github.com/guilherme-luvi/go-api-fiber-crud/src/schemas"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{db}
}

func (u *user) CreateUser(user schemas.User) error {
	return u.db.Create(&user).Error
}

func (u *user) GetUsers() ([]schemas.User, error) {
	var users []schemas.User
	err := u.db.Find(&users).Error
	return users, err
}
