//go:build wireinject
// +build wireinject

package cmd

import (
	"bookstack/repository"
	"bookstack/router"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func newServer(db *gorm.DB, repo repository.Repository) (*Server, error) {
	wire.Build(
		router.Setup,
		wire.Struct(new(Server), "*"),
	)
	return nil, nil
}
