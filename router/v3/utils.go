package v3

import (
	"bookstack/router/extension"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotImplemented() error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

func bindAndValidate(c echo.Context, i interface{}) error {
	return extension.BindAndValidate(c, i)
}
