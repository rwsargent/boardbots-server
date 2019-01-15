package context

import (
	"github.com/labstack/echo"
	"github.com/google/uuid"
	"boardbots/quoridor"
)

type
(
	BBContext interface {
		GetPlayerName() string
		GetPlayerId()
	}

	DefaultBBContext struct {
		echo.Context
		PlayerPrinciple PlayerPrinciple
		Game *quoridor.Game
	}

	PlayerPrinciple struct {
		UserName string
		UserId uuid.UUID
		Password string
	}
)

func (ctx *DefaultBBContext) GetPlayerName() string {
	return ctx.PlayerPrinciple.UserName
}

func (ctx *DefaultBBContext) GetPlayerId() uuid.UUID {
	return ctx.PlayerPrinciple.UserId
}
