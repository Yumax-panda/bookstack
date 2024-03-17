package cmd

import (
	"bookstack/repository"

	"github.com/labstack/echo"
)

type Server struct {
	Router *echo.Echo
	Repo   repository.Repository
}

func (s *Server) Start(address string) error {
	return s.Router.Start(address)
}
