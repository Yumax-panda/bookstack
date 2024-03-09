package migration

import (
	"bookstack/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
