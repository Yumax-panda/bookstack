package gorm

import (
	"gorm.io/gorm"

	"bookstack/repository"
)

func convertError(err error) error {
	switch {
	case err == gorm.ErrRecordNotFound:
		return repository.ErrNotFound
	default:
		return err
	}
}
