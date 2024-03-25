package v3

import (
	"bookstack/repository"
	"bookstack/router/middlewares"
	"bookstack/router/session"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Repo      repository.Repository
	SessStore session.Store
}

func (h *Handlers) Setup(e *echo.Group) {
	noLogin := middlewares.NoLogin(h.SessStore, h.Repo)

	apiNoAuth := e.Group("/v3")
	{
		apiNoAuth.POST("/users", h.CreateUser, noLogin)
	}
}
