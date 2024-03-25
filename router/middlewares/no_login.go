package middlewares

import (
	"bookstack/repository"
	"bookstack/router/extension/herror"
	"bookstack/router/session"

	"github.com/labstack/echo/v4"
)

// NoLogin セッションが既に存在するリクエストを拒否するミドルウェア

func NoLogin(sessStore session.Store, repo repository.Repository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if len(c.Request().Header.Get(echo.HeaderAuthorization)) > 0 {
				return herror.BadRequest("Authorization Header must not be set. Please logout once.")
			}

			sess, err := sessStore.GetSession(c)

			if err != nil {
				return herror.InternalServerError(err)
			}

			if sess != nil && sess.LoggedIn() {
				_, err := repo.GetUser(sess.UserID(), false)

				if err != nil {
					return herror.InternalServerError(err)
				}

				return herror.BadRequest("You are already logged in. Please logout once.")
			}

			return next(c)
		}
	}
}
