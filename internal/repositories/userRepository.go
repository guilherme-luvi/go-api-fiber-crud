package repositories

import (
	"github.com/guilherme-luvi/go-api-fiber-crud/internal/schemas"
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

func (u *user) GetAllUsers() ([]schemas.User, error) {
	var users []schemas.User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *user) GetUserByID(id string, user *schemas.User) error {
	return u.db.First(user, id).Error
}

func (u *user) GetUserByEmail(email string, user *schemas.User) error {
	return u.db.Where("email = ?", email).First(user).Error
}

func (u *user) UpdateUser(id string, user schemas.User) error {
	return u.db.Model(&schemas.User{}).Where("id = ?", id).Updates(&user).Error
}

func (u *user) DeleteUser(id string) error {
	return u.db.Delete(&schemas.User{}, id).Error
}
