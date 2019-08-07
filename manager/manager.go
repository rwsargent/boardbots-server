package manager

import (
	q "boardbots-server/quoridor"
	"errors"
	"github.com/google/uuid"
)

const (
	GameState_Started  = "started"
	GameState_Finished = "finished"
)

type (
	GameParameters struct {
		GameState string
		Players   []string
	}

	GameManager interface {
		GetGame(gameId uuid.UUID) (ManagedGame, error)
		AddGame(game q.Game) error
		GetGamesForUser(playerId uuid.UUID) ([]ManagedGame, error)
		FindGames(params GameParameters) ([]ManagedGame, error)
	}

	Filter func(game *q.Game) bool

	ManagedGame struct {
		modCount int
		Game     q.Game
	}
)

// For use in testingutils
type FakeMemoryManager struct {
	ReturnId    uuid.UUID
	ReturnGame  *q.Game
	ReturnError error
}

func (manager *FakeMemoryManager) GetGame(gameId uuid.UUID) (*q.Game, error) {
	if manager.ReturnGame == nil {
		return nil, errors.New("Error!")
	}
	return manager.ReturnGame, nil
}

func (manager *FakeMemoryManager) AddGame(game *q.Game) error {
	return nil
}

func (manager *FakeMemoryManager) GetGamesForUser(playerId uuid.UUID) ([]q.Game, error) {
	return nil, nil
}

func (manager *FakeMemoryManager) FindGames(params GameParameters) ([]q.Game, error) {
	return nil, nil
}
