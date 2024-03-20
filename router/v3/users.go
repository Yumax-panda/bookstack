package v3

import (
	"bookstack/repository"
	"bookstack/router/extension"
	"bookstack/router/extension/herror"

	"github.com/labstack/echo/v4"
)

// GetUsers GET /users
func (h *Handlers) GetUsers(c echo.Context) error {
	q := repository.UsersQuery{}

	name := c.QueryParam("name")

	if len(name) > 0 {
		q = q.NameOf(name)
	}

	users, err := h.Repo.GetUsers(q)

	if err != nil {
		return herror.InternalServerError(err)
	}

	return extension.ServeJSONWithETag(c, formatUsers(users))

}
