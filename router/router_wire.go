//go:build wireinject
// +build wireinject

package router

import (
	"bookstack/repository"
	"bookstack/router/session"
	v3 "bookstack/router/v3"

	"github.com/google/wire"

	"gorm.io/gorm"
)

func newRouter(db *gorm.DB, repo repository.Repository) *Router {
	wire.Build(
		newEcho,
		session.NewGormStore,
		wire.Struct(new(v3.Handlers), "*"),
		wire.Struct(new(Router), "*"),
	)
	return nil
}
