package middleware

import (
	"github.com/labstack/echo"
	"boardbots/manager"
	"boardbots/server/context"
	"boardbots/server/transport"
	"boardbots/quoridor"
	"github.com/google/uuid"
	"net/http"
)

type (
	GameAuthenticator struct {
		GameManager manager.GameManager
	}
)
func (ga *GameAuthenticator) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var gameRequest transport.GameRequest
		if err := ctx.Bind(&gameRequest); err != nil {
			return transport.StandardBadRequestError(err)
		}
		game, err := ga.GameManager.GetGame(gameRequest.GameId)
		if err != nil {
			return transport.StandardBadRequestError(err)
		}
		bbCtx := ctx.(context.DefaultBBContext)
		if !playerInGame(game.Players, bbCtx.PlayerPrinciple.UserId) {
			return echo.NewHTTPError(http.StatusForbidden, transport.BaseResponse{
				Error: "you do not have access to this game",
			})
		}
		bbCtx.Game = game
		return next(ctx)
	}
}

func playerInGame(players map[quoridor.PlayerPosition]*quoridor.Player, playerId uuid.UUID) bool {
	for _, player := range players {
		if player.PlayerId == playerId {
			return true
		}
	}
	return false
}