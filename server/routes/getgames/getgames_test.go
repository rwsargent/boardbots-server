package getgames

import (
	"testing"
	tu "boardbots-server/server/testingutils"
	"boardbots-server/server/mocks"
	"boardbots-server/manager"
	"github.com/stretchr/testify/assert"
	q "boardbots-server/quoridor"
	"github.com/stretchr/testify/mock"
	"errors"
	"boardbots-server/server/transport"
	"boardbots-server/server/context"
)

func TestGetGames_HappyPath(t *testing.T) {
	mockManager := new(mocks.MockGameManager)
	mockManager.On("FindGames", mock.Anything).
		Return([]manager.ManagedGame{
			{Game: *q.NewTwoPersonGame(tu.TestUUID)},
			{Game:*q.NewTwoPersonGame(tu.TestMissingUUID)},
			{Game: *q.NewTwoPersonGame(tu.GameIds[0])},}, nil)
	bbCtx, recorder := context.Build(context.DefaultFakeContextBuilder().Override(context.FakeContextBuilder{
		Payload: "{}",
	}))
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
	mockManager.On("FindGames", mock.Anything).
		Return([]q.Game{}, nil)
	bbCtx, recorder := context.Build(context.DefaultFakeContextBuilder())
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
	bbCtx, _ := context.Build(context.DefaultFakeContextBuilder())
	h := fakeHandler(mockManager)

	result := h.Handle(bbCtx)
	if assert.Error(t, result) {
		httpError := result.(transport.HandledError)
		assert.Equal(t, httpError.Message, "error")
	}
}

func TestOverride(t *testing.T) {
	base := context.FakeContextBuilder{
		Payload : "base",
		Path : "base",
		Game: *q.NewTwoPersonGame(tu.GameIds[0]),
		Method: "base",
		Headers: nil,
	}

	override := base.Override(context.FakeContextBuilder{
		Payload:"override",
		Method : "override",
		Game : q.Game{},
	})

	assert.Equal(t, override.Payload, "override")
	assert.Equal(t, "base", override.Path)
	assert.Equal(t, override.Method, "override")
	assert.Nil(t, override.Headers)
	assert.Equal(t, override.Game, base.Game)


}

func fakeHandler(gameManager manager.GameManager) Handler {
	return Handler{
		GameManager: gameManager,
	}
}
