package router

import (
	"bookstack/repository"
	v3 "bookstack/router/v3"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Router struct {
	e  *echo.Echo
	v3 *v3.Handlers
}

func Setup(db *gorm.DB, repo repository.Repository) *echo.Echo {
	r := newRouter(db, repo)

	api := r.e.Group("/api")
	r.v3.Setup(api)

	return r.e
}

func newEcho(repo repository.Repository) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// TODO: Add middleware

	return e
}
