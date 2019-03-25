package context

import (
	"boardbots/quoridor"
	"boardbots/server/player"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type (
	BBContext interface {
		GetPlayerName() string
		GetPlayerId() uuid.UUID
	}

	DefaultBBContext struct {
		echo.Context
		PlayerPrinciple player.PlayerPrinciple
		Game            *quoridor.Game
	}
)

func (ctx *DefaultBBContext) GetPlayerName() string {
	return ctx.PlayerPrinciple.UserName
}

func (ctx *DefaultBBContext) GetPlayerId() uuid.UUID {
	return ctx.PlayerPrinciple.UserId
}
