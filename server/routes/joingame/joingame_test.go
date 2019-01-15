package joingame

import (
	"testing"

	"github.com/labstack/echo"
	"net/http/httptest"
	"strings"
	"net/http"
	"boardbots/manager"
	"github.com/google/uuid"
	"boardbots/quoridor"
	"github.com/stretchr/testify/assert"
	"boardbots/server/transport"
	"boardbots/server/context"
	"encoding/json"
	"boardbots/server/testingutils"
)

func TestHandler_JoinGame(t *testing.T) {
	ctx, recorder := fakeContext(http.MethodPost, "/joingame", `{"gameId":"2ce1f41c-3a2e-402a-852b-737321e3ec7d"}`)
	bbCtx := context.DefaultBBContext{ctx, context.PlayerPrinciple{UserName: "name"}, nil}
	handler := fakeHandler("12345432-3a2e-402a-852b-737321e3ec7d", &quoridor.Game{})

	result := handler.JoinGame(bbCtx)

	if assert.NoError(t, result) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		response := &transport.BaseResponse{}
		json.NewDecoder(recorder.Body).Decode(response)
		assert.Equal(t, transport.BaseResponse{}, *response)
	}
}

func TestUnknownGame(t *testing.T) {
	ctx, _ := fakeContext(http.MethodPost, "/joingame", `{"gameId":"2ce1f41c-3a2e-402a-852b-737321e3ec7d"}`)
	bbCtx := context.DefaultBBContext{ctx, context.PlayerPrinciple{UserName: "name"}, nil}
	handler := fakeHandler("12345432-3a2e-402a-852b-737321e3ec7d", nil)

	result := handler.JoinGame(bbCtx).(*echo.HTTPError)
	if assert.Error(t, result) {
		assert.Equal(t, http.StatusBadRequest, result.Code)
	}
}

func TestPlayerIsAddedToGame(t *testing.T) {
	ctx, recorder := fakeContext(http.MethodPost, "/joingame", `{"gameId":"2ce1f41c-3a2e-402a-852b-737321e3ec7d"}`)
	bbCtx := context.DefaultBBContext{
		ctx,
		context.PlayerPrinciple{
			UserName: "name",
			UserId: uuid.MustParse("54321543-3a2e-402a-852b-737321e3ec7d")},
		nil}
	game := quoridor.NewTwoPersonGame()
	game.Players[quoridor.PlayerOne].PlayerId = uuid.New()
	handler := fakeHandler("12345432-3a2e-402a-852b-737321e3ec7d", game)

	result := handler.JoinGame(bbCtx)

	if assert.NoError(t, result) {
		var response Response
		testingutils.FillResponseFromPayload(recorder, &response)
		assert.Equal(t, Response{PlayerPosition: quoridor.PlayerTwo}, response)
		assert.Len(t, game.Players, 2)
		assert.NotEqual(t, uuid.Nil, game.Players[quoridor.PlayerTwo].PlayerId)
	}
}

func fakeContext(method, path, payload string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	return e.NewContext(req, recorder), recorder
}

func fakeHandler(id string, game *quoridor.Game) Handler {
	gameManager := manager.FakeMemoryManager{ReturnId: uuid.MustParse(id), ReturnGame: game}
	return Handler{GameManager: &gameManager}
}
