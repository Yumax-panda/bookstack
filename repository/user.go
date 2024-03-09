package repository

import "bookstack/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
}
