package migration

import (
	"bookstack/model"

	"gorm.io/gorm"
)

func AllTables() []interface{} {
	return []interface{}{
		&model.User{},
		&model.UserProfile{},
		&model.Note{},
		&model.SessionRecord{},
	}
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(AllTables()...)
}
