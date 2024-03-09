package gorm

import (
	"bookstack/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// FindAll implements UserRepository interface
func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
