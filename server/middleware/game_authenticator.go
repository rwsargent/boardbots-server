package middleware

import (
	"boardbots-server/manager"
	"boardbots-server/quoridor"
	"boardbots-server/server/context"
	"boardbots-server/server/transport"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

type (
	GameAuthenticator struct {
		GameManager manager.GameManager
	}
)

func Authenticator(gameManager manager.GameManager) func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	mid := GameAuthenticator{
		GameManager: gameManager,
	}
	return mid.Authenticate
}

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
		if !playerInGame(game.Game.Players, bbCtx.PlayerPrinciple.UserId) {
			return echo.NewHTTPError(http.StatusForbidden, transport.BaseResponse{
				Error: "you do not have access to this game",
			})
		}
		bbCtx.Game = &game.Game
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
