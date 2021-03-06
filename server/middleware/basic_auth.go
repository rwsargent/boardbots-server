package middleware

import (
	"boardbots-server/server/context"
	"boardbots-server/server/persistence"
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func GetBasicAuthenticator(userPortal persistence.UserPortal) middleware.BasicAuthConfig {
	return middleware.BasicAuthConfig{
		Skipper: middleware.DefaultSkipper,
		Validator: func(username, password string, ctx echo.Context) (bool, error) {
			if userPortal.ValidateCredentials(username, password) {
				bbCtx := ctx.(*context.DefaultBBContext)
				principle, err := userPortal.GetPlayerPrinciple(username)
				if err == nil {
					bbCtx.PlayerPrinciple = principle
					return true, nil
				} else {
					return false, err
				}
			} else {
				return false, errors.New("invalid credentials")
			}
		},
		Realm: "Restricted",
	}
}
