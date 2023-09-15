package config

import (
	"gorm.io/gorm"
	"movieProject/data/Entity"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&Entity.Movie{})
}
