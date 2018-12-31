package manager

import (
	"github.com/google/uuid"
	q "boardbots/quoridor"
	"errors"
	"fmt"
)

type (
	GameManager interface {
		GetGame(gameId uuid.UUID) (*q.Game, error)
		AddGame(game *q.Game) (uuid.UUID, error)
	}

	InMemoryGameManager struct {
		games map[uuid.UUID]*q.Game
	}
)
func (manager *InMemoryGameManager) CreateTwoPlayerGame() *q.Game{
	game := q.NewTwoPersonGame()
	gid := uuid.New()
	manager.games[gid] = game
	return game
}

func (manager *InMemoryGameManager) AddGame(game *q.Game) (uuid.UUID, error) {
	gameId, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}
	manager.games[gameId] = game
	return gameId, nil
}

func (manager *InMemoryGameManager) GetGame(gameId uuid.UUID) (*q.Game, error) {
	game, ok := manager.games[gameId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no game found with id: %s", gameId))
	}
	return game, nil
}

func NewMemoryGameManager() *InMemoryGameManager{
	return &InMemoryGameManager{games : make(map[uuid.UUID]*q.Game)}
}

// For use in testingutils
type FakeMemoryManager struct {
	ReturnId uuid.UUID
	ReturnGame *q.Game
	ReturnError error
}

func (manager *FakeMemoryManager) GetGame(gameId uuid.UUID) (*q.Game, error) {
	if manager.ReturnGame == nil {
		return nil, errors.New("Error!")
	}
	return manager.ReturnGame, nil
}

func (manager *FakeMemoryManager) AddGame(game *q.Game) (uuid.UUID, error) {
	return manager.ReturnId, nil;
}