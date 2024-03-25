package extension

import (
	"bookstack/router/extension/herror"

	vd "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/labstack/echo/v4"

	"bookstack/router/utils"
)

func BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := vd.ValidateWithContext(utils.NewRequestValidateContext(c), i); err != nil {
		if e, ok := err.(vd.InternalError); ok {
			return herror.InternalServerError(e.InternalError())
		}
		return herror.BadRequest(err)
	}
	return nil
}
