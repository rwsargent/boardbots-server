package getgame

import (
	"testing"
	tu "boardbots/server/testingutils"
	"net/http"
	"boardbots/server/context"
	"boardbots/manager"
	"github.com/google/uuid"
	"boardbots/quoridor"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo"
	"boardbots/server/transport"
)

func TestHandler_GetGame_WithUnparsableId(t *testing.T) {
 ctx, _ := tu.FakeContext(
 	http.MethodGet,
 	"/getgame?gameid=invalid",
"")

 bbCtx := context.DefaultBBContext{ctx,context.PlayerPrinciple{UserName:"name"}}
 handler := fakeHandler(tu.TestUUID, quoridor.NewTwoPersonGame())

 result := handler.GetGame(bbCtx).(*echo.HTTPError)
 assert.Error(t, result)
 assert.Equal(t, "could not parse id: invalid", result.Message.(transport.BaseResponse).Error)
 assert.Equal(t, http.StatusBadRequest, result.Code)
}

func fakeHandler(id uuid.UUID, game *quoridor.Game) Handler {
	gameManager := manager.FakeMemoryManager{ReturnId: id, ReturnGame: game}
	return Handler{GameManager: &gameManager}
}

func TestHandle_GetGame_MissingQuery(t *testing.T) {
	ctx, _ := tu.FakeContext(http.MethodGet,"/getgame","")
	bbCtx := context.DefaultBBContext{ctx,context.PlayerPrinciple{UserName:"name"}}
	handler := fakeHandler(tu.TestUUID, quoridor.NewTwoPersonGame())

	result := handler.GetGame(bbCtx).(*echo.HTTPError)

	assert.Error(t, result)
	assert.Equal(t, "no gameid query parameter", result.Message.(transport.BaseResponse).Error)
	assert.Equal(t, http.StatusBadRequest, result.Code)
}

func TestHandle_GetGame_FindsValidGame(t *testing.T) {
	ctx, recorder := tu.FakeContext(http.MethodGet,"/getgame?gameid=" + tu.TestUUID.String(),"")
	bbCtx := context.DefaultBBContext{ctx,context.PlayerPrinciple{UserName:"name"}}
	handler := fakeHandler(tu.TestUUID, quoridor.NewTwoPersonGame())

	result := handler.GetGame(bbCtx)
	var response transport.GameResponse
	tu.FillResponseFromPayload(recorder, &response)
	assert.NoError(t, result)
	assert.Len(t, response.Board, 2)
	assert.Equal(t, quoridor.PlayerOne, response.CurrentTurn)
}

func TestHandle_GetGame_FindsInvalidGame(t *testing.T) {
	ctx, _ := tu.FakeContext(http.MethodGet,"/getgame?gameid=" + tu.TestMissingUUID.String(),"")
	bbCtx := context.DefaultBBContext{ctx,context.PlayerPrinciple{UserName:"name"}}
	handler := fakeHandler(tu.TestUUID, nil)

	result := handler.GetGame(bbCtx).(*echo.HTTPError)

	assert.Error(t, result)
	assert.Equal(t, http.StatusBadRequest, result.Code)
}