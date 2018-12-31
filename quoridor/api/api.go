package api

import (
	"boardbots/manager"
	"boardbots/quoridor"
)

type Api struct {
	GameManager manager.GameManager
}

func (api *Api) MovePawn(command BaseCommand) (quoridor.Board, error) {
	game, err := api.GameManager.GetGame(command.GameId)
	if err != nil {
		return nil, err
	}
	game.MovePawn(command.Position, )

}