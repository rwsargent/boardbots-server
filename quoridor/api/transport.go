package api

import (
	"github.com/google/uuid"
	"boardbots/util"
)

type BaseCommand struct {
	GameId uuid.UUID `json: gameId`
	Position util.Position `json: position`
}

