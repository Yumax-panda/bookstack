package utils

import (
	"bookstack/router/consts"
	"context"

	"github.com/labstack/echo/v4"
)

type ctxKey int

const (
	repoctxKey ctxKey = iota
	cmctxKey
)

func NewRequestValidateContext(c echo.Context) context.Context {
	return context.WithValue(context.WithValue(context.Background(), repoctxKey, c.Get(consts.KeyRepo)), cmctxKey, c.Get(consts.KeyChannelManager))
}
