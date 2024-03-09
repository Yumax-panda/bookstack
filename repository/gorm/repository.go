package gorm

import (
	"bookstack/repository"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
	repository.UserRepository
}
