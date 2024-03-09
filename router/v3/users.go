package v3

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handlers) GetAllUsers(c echo.Context) error {
	users, err := h.Repo.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
