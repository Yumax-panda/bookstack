package middlewares

import (
	"bookstack/repository"
	"bookstack/router/consts"
	"bookstack/router/extension/ctxkey"
	"bookstack/router/extension/herror"
	"bookstack/router/session"
	"context"

	"github.com/labstack/echo/v4"
)

// UserAuthenticate ユーザー認証ミドルウェア
// ユーザーがログインしているかを確認し、ログインしている場合はユーザー情報をコンテキストにセットする。
func UserAuthenticate(repo repository.Repository, sessStore session.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := sessStore.GetSession(c)
			if err != nil {
				return herror.InternalServerError(err)
			}

			if sess == nil || !sess.LoggedIn() {
				return herror.Unauthorized("You must be logged in to access this resource.")
			}

			uid := sess.UserID()

			user, err := repo.GetUser(uid, true)
			if err != nil {
				return herror.InternalServerError(err)
			}

			c.Set(consts.KeyUser, user)
			c.Set(consts.KeyUserID, user.GetID())
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), ctxkey.UserID, user.GetID())))

			return next(c)
		}
	}
}
