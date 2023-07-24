package database

import (
	"api.virak.me/modules/blog/domain/entities"
	userEntity "api.virak.me/modules/user/domain/entities"
	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) {
	db.AutoMigrate(&entities.BlogEntity{})
	db.AutoMigrate(&userEntity.UserEntity{})
}
