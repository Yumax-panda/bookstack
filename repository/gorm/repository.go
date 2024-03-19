package gorm

import (
	"bookstack/migration"
	"bookstack/repository"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
	repository.UserRepository
	repository.NoteRepository
}

func NewGormRepository(db *gorm.DB) repository.Repository {
	repo := &Repository{
		DB:             db,
		UserRepository: makeUserRepository(db),
		NoteRepository: makeNoteRepository(db),
	}
	migration.Migrate(db)
	fmt.Println("migration was successful")
	return repo
}
