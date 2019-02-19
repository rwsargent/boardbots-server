package getgames

import (
	"testing"
	tu "boardbots/server/testingutils"
	"boardbots/server/mocks"
	"boardbots/manager"
	"github.com/stretchr/testify/assert"
	q "boardbots/quoridor"
	"github.com/stretchr/testify/mock"
	"errors"
	"boardbots/server/transport"
)

func TestGetGames_HappyPath(t *testing.T) {
	mockManager := new(mocks.MockGameManager)
	mockManager.On("GetGamesForUser", mock.Anything).
		Return([]q.Game{*q.NewTwoPersonGame(), *q.NewTwoPersonGame(), *q.NewTwoPersonGame()}, nil)
	bbCtx, recorder := tu.Build(tu.DefaultFakeContextBuilder())
	h := fakeHandler(mockManager)

	result := h.Handle(bbCtx)
	if assert.NoError(t, result) {
		body := &Response{}
		tu.ReadBodyFromRecorder(recorder, body)
		assert.Len(t, body.Games, 3)
	}
}

func TestGetGames_NoErrorNoGames(t *testing.T) {
	mockManager := new(mocks.MockGameManager)
	mockManager.On("GetGamesForUser", mock.Anything).
		Return([]q.Game{}, nil)
	bbCtx, recorder := tu.Build(tu.DefaultFakeContextBuilder())
	h := fakeHandler(mockManager)

	result := h.Handle(bbCtx)
	if assert.NoError(t, result) {
		body := &Response{}
		tu.ReadBodyFromRecorder(recorder, body)
		assert.Len(t, body.Games, 0)
	}
}

func TestGetGames_ErrorInGameManager(t *testing.T) {
	mockManager := new(mocks.MockGameManager)
	mockManager.On("GetGamesForUser", mock.Anything).
		Return(nil, errors.New("error"))
	bbCtx, _ := tu.Build(tu.DefaultFakeContextBuilder())
	h := fakeHandler(mockManager)

	result := h.Handle(bbCtx)
	if assert.Error(t, result) {
		httpError := result.(transport.HandledError)
		assert.Equal(t, httpError.Message, "error")
	}
}

func fakeHandler(gameManager manager.GameManager) Handler {
	return Handler{
		GameManager: gameManager,
	}
}