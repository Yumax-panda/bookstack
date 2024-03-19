package v3

import (
	"bookstack/repository"

	"github.com/labstack/echo"
)

type Handlers struct {
	Repo repository.Repository
}

func (h *Handlers) Setup(e *echo.Group) {

}
