package database

import (
	"api.virak.me/src/modules/blog/domain/entities"
	userEntity "api.virak.me/src/modules/user/domain/entities"
	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) {
	db.AutoMigrate(&entities.BlogEntity{})
	db.AutoMigrate(&userEntity.UserEntity{})
}
