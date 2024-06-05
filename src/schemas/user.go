package schemas

import "gorm.io/gorm"

// User schema
type User struct {
	gorm.Model

	Name  string `gorm:"not null"`
	Email string `gorm:"not null,unique"`
}
