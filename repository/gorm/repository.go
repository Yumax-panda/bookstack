package gorm

import (
	"bookstack/repository"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
	repository.UserRepository
}

func NewGormRepository(db *gorm.DB) repository.Repository {
	return &Repository{
		DB:             db,
		UserRepository: makeUserRepository(db),
	}
}
