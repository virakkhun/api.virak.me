package entities

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Role     string
	Photo    string
}
