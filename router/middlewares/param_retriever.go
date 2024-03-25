package middlewares

import (
	"bookstack/repository"
	"bookstack/router/consts"
	"bookstack/router/extension/herror"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

type ParamRetriever struct {
	repo repository.Repository
}

func NewParamRetriever(repo repository.Repository) *ParamRetriever {
	return &ParamRetriever{repo: repo}
}

func (pr *ParamRetriever) byString(param string, key string, f func(c echo.Context, v string) (interface{}, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r, err := f(c, c.Param(param))
			if err != nil {
				return pr.error(err)
			}
			c.Set(key, r)
			return next(c)
		}
	}
}

func (pr *ParamRetriever) byUUID(param string, key string, f func(e echo.Context, v uuid.UUID) (interface{}, error)) echo.MiddlewareFunc {
	return pr.byString(param, key, func(c echo.Context, v string) (interface{}, error) {
		u, err := uuid.FromString(v)
		if err != nil || u == uuid.Nil {
			return nil, herror.NotFound()
		}
		return f(c, u)
	})
}

func (pr *ParamRetriever) checkOnlyByUUID(param string, f func(c echo.Context, v uuid.UUID) (bool, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			v, err := uuid.FromString(c.Param(param))
			if err != nil {
				return herror.NotFound()
			}
			ok, err := f(c, v)
			if err != nil {
				return pr.error(err)
			}
			if !ok {
				return herror.NotFound()
			}

			return next(c)
		}
	}
}

func (pr *ParamRetriever) error(err error) error {
	switch err.(type) {
	case *echo.HTTPError:
		return err
	case *herror.InternalError:
		return err
	default:
		if err == repository.ErrNotFound {
			return herror.NotFound()
		}
		return herror.InternalServerError(err)
	}
}

func (pr *ParamRetriever) UserID(checkOnly bool) echo.MiddlewareFunc {
	if checkOnly {
		return pr.checkOnlyByUUID(consts.ParamUserID, func(_ echo.Context, v uuid.UUID) (bool, error) {
			return pr.repo.UserExists(v)
		})
	}
	return pr.byUUID(consts.ParamUserID, consts.KeyParamUser, func(_ echo.Context, v uuid.UUID) (interface{}, error) {
		return pr.repo.GetUser(v, true)
	})
}
