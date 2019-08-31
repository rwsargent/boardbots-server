package manager

import (
	q "boardbots-server/quoridor"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"sync"
)

type (
	InMemoryGameManager struct {
		games map[uuid.UUID]ManagedGame
		locks map[uuid.UUID]sync.RWMutex
	}
)

func (manager *InMemoryGameManager) FindGames(params GameParameters) ([]ManagedGame, error) {
	panic("implement me")
}

func (manager *InMemoryGameManager) AddGame(game q.Game) error {
	gid := game.Id
	if gid == uuid.Nil {
		return errors.New("game id cannot be nil")
	}
	if _, present := manager.games[gid]; present {
		return errors.New(fmt.Sprintf("game %s is already managed", gid))
	}
	manager.games[gid] = ManagedGame{Game: game.Copy()}
	manager.locks[gid] = sync.RWMutex{}
	return nil
}

func (manager *InMemoryGameManager) GetGamesForUser(playerId uuid.UUID) ([]ManagedGame, error) {
	games := make([]ManagedGame, 0)
	for _, game := range manager.games {
		for _, player := range game.Game.Players {
			if player.PlayerId == playerId {
				games = append(games, game)
			}
		}
	}
	return games, nil
}

func (manager *InMemoryGameManager) GetGame(gameId uuid.UUID) (ManagedGame, error) {
	lock, ok := manager.locks[gameId]
	if !ok {
		return ManagedGame{}, errors.New(fmt.Sprintf("no game found with id: %s", gameId))
	}
	lock.RLock()
	defer lock.RUnlock()
	game, ok := manager.games[gameId]
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
			lock := manager.locks[game.Game.Id]
			lock.RLock()
			shouldAdd = shouldAdd || filter(&game.Game)
			lock.RUnlock()
		}
		if shouldAdd {
			games = append(games, ManagedGame{Game: game.Game.Copy(), modCount: game.modCount})
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
	return &InMemoryGameManager{
		games: make(map[uuid.UUID]ManagedGame),
		locks: make(map[uuid.UUID]sync.RWMutex),
	}
}
