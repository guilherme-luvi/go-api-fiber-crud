package schemas

import "gorm.io/gorm"

// User schema
type User struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
