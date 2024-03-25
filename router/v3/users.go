package v3

import (
	"bookstack/repository"
	"bookstack/router/extension"
	"bookstack/router/extension/herror"
	"bookstack/utils/validator"
	"net/http"

	vd "github.com/go-ozzo/ozzo-validation/v4"

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

type PostUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r *PostUserRequest) Validate() error {
	return vd.ValidateStruct(&r,
		vd.Field(&r.Name, validator.UserNameRuleRequired...),
		vd.Field(&r.Password, validator.PasswordRuleRequired...),
	)
}

func (h *Handlers) CreateUser(c echo.Context) error {
	var req PostUserRequest

	if err := bindAndValidate(c, &req); err != nil {
		return err
	}

	user, err := h.Repo.CreateUser(repository.CreateUserArgs{
		Name:     req.Name,
		Password: req.Password,
	})

	if err != nil {
		switch err {
		case repository.ErrAlreadyExists:
			return herror.Conflict("name conflicts")
		default:
			return herror.InternalServerError(err)
		}
	}

	return c.JSON(http.StatusCreated, formatUserDetail(user))
}
