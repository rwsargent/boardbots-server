package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/google/uuid"
	q "boardbots/quoridor"
)

type MockGameManager struct {
	mock.Mock
}

func (mock *MockGameManager) GetGame(gameId uuid.UUID) (*q.Game, error){
	args := mock.Called(gameId)
	return args.Get(0).(*q.Game), args.Error(1)
}

func (mock *MockGameManager) AddGame(game *q.Game) (uuid.UUID, error){
	args := mock.Called(game)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (mock *MockGameManager) GetGamesForUser(playerId uuid.UUID) ([]q.Game, error){
	args := mock.Called(playerId)
	if args.Get(0) != nil {
		return args.Get(0).([]q.Game), args.Error(1)
	} else  {
		return nil, args.Error(1)
	}
}

