package mocks

import (
	"boardbots-server/manager"
	q "boardbots-server/quoridor"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockGameManager struct {
	mock.Mock
}

func (mock *MockGameManager) GetGame(gameId uuid.UUID) (manager.ManagedGame, error) {
	args := mock.Called(gameId)
	return args.Get(0).(manager.ManagedGame), args.Error(1)
}

func (mock *MockGameManager) AddGame(game q.Game) error {
	args := mock.Called(game)
	return args.Error(1)
}

func (mock *MockGameManager) GetGamesForUser(playerId uuid.UUID) ([]manager.ManagedGame, error) {
	args := mock.Called(playerId)
	if args.Get(0) != nil {
		return args.Get(0).([]manager.ManagedGame), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (mock *MockGameManager) FindGames(parameters manager.GameParameters) ([]manager.ManagedGame, error) {
	args := mock.Called(parameters)
	if args.Get(0) != nil {
		return args.Get(0).([]manager.ManagedGame), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}
