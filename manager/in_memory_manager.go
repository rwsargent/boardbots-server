package manager

import (
	q "boardbots/quoridor"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type (
	InMemoryGameManager struct {
		games map[uuid.UUID]ManagedGame
	}
)

func (manager *InMemoryGameManager) AddGame(game q.Game) error {
	gid := game.Id
	if gid == uuid.Nil {
		return errors.New("game id cannot be nil")
	}
	if _, present := manager.games[gid]; present {
		return errors.New(fmt.Sprintf("game %s is already managed", gid))
	}
	manager.games[gid] = ManagedGame{Game: game.Copy()}
	return nil
}

func (manager *InMemoryGameManager) GetGamesForUser(playerId uuid.UUID) ([]q.Game, error) {
	games := make([]q.Game, 0)
	for _, game := range manager.games {
		for _, player := range game.Game.Players {
			if player.PlayerId == playerId {
				games = append(games, game.Game)
			}
		}
	}
	return games, nil
}

func (manager *InMemoryGameManager) GetGame(gameId uuid.UUID) (ManagedGame, error) {
	game, ok := manager.games[gameId]
	if !ok {
		return ManagedGame{}, errors.New(fmt.Sprintf("no game found with id: %s", gameId))
	}
	game.Game = game.Game.Copy()
	return game, nil
}

func (manager *InMemoryGameManager) FindGame(params GameParameters) ([]ManagedGame, error) {
	filters, err := filtersFromParameters(params)
	if err != nil {
		return nil, err
	}
	games := make([]ManagedGame, 0)
	for _, game := range manager.games {
		var shouldAdd bool
		for _, filter := range filters {
			shouldAdd = shouldAdd || filter(&game.Game)
		}
		if shouldAdd {
			games = append(games, game)
		}
	}
	return games, nil
}

func filtersFromParameters(parameters GameParameters) ([]Filter, error) {
	filters := make([]Filter, 0)
	switch strings.ToLower(parameters.GameState) {
	case "open":
		filters = append(filters, func(game *q.Game) bool {
			return game.StartDate.IsZero()
		})
	case "started":
		filters = append(filters, func(game *q.Game) bool {
			return !game.StartDate.IsZero() && game.EndDate.IsZero()
		})
	case "finished":
		filters = append(filters, func(game *q.Game) bool {
			return !game.StartDate.IsZero() && !game.EndDate.IsZero()
		})
	}

	if len(parameters.Players) != 0 {
		playerIds := make(map[uuid.UUID]bool)
		for _, player := range parameters.Players {
			playerId, err := uuid.Parse(player)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("playerId %s could not be parsed as an id", player))
			}
			playerIds[playerId] = true
		}

		filters = append(filters, func(game *q.Game) bool {
			for _, player := range game.Players {
				if playerIds[player.PlayerId] {
					return true
				}
			}
			return false
		})
	}
	return filters, nil
}

func NewMemoryGameManager() *InMemoryGameManager {
	return &InMemoryGameManager{games: make(map[uuid.UUID]ManagedGame)}
}
