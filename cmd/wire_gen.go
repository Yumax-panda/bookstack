// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"bookstack/repository"
	"bookstack/router"
	"gorm.io/gorm"
)

// Injectors from serve_wire.go:

func newServer(db *gorm.DB, repo repository.Repository) (*Server, error) {
	echo := router.Setup(db, repo)
	server := &Server{
		Router: echo,
		Repo:   repo,
	}
	return server, nil
}