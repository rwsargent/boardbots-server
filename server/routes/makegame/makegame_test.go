package makegame

import (
	"testing"

	"net/http"
	"github.com/labstack/echo"
	"net/http/httptest"
	"strings"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"boardbots/server/context"
	"encoding/json"
	"boardbots/manager"
)

func TestHandler_MakeGame(t *testing.T) {
	ctx, recorder := fakeContext(http.MethodPost, "/makegame", `{"PlayerCount" : 2}`)
	id := uuid.MustParse("2ce1f41c-3a2e-402a-852b-737321e3ec7d")
	bbCtx := context.DefaultBBContext{ctx, context.PlayerPrinciple{"name", id}}
	h := fakeHandler("2ce1f41c-3a2e-402a-852b-737321e3db7a")
	result := h.MakeGame(bbCtx)

	if assert.NoError(t, result, "expecting successful run") {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.NotEmpty(t, h)
		res := getPayloadFromResponse(recorder)
		assert.Equal(t, res.GameId, uuid.MustParse("2ce1f41c-3a2e-402a-852b-737321e3db7a"))
	}
}

func TestHandler_MakeGame_NoPlayerPrinciple(t *testing.T) {
	ctx, _ := fakeContext(http.MethodPost, "/makegame", `{"PlayerCount" : 2}`)
	bbCtx := context.DefaultBBContext{ctx, context.PlayerPrinciple{}}
	h := fakeHandler("2ce1f41c-3a2e-402a-852b-737321e3db7a")

	result := h.MakeGame(bbCtx).(*echo.HTTPError)

	payload := getPayloadFromResult(result)
	if assert.Error(t, result, "expecting successful run") {
		assert.Equal(t, http.StatusInternalServerError, result.Code)
		assert.Equal(t, "expected a player name, didn't find one", payload.Error)
	}
}
func getPayloadFromResult(httpError *echo.HTTPError) Response {
	return httpError.Message.(Response)
}

func fakeHandler(id string) Handler {
	gameManager := manager.FakeMemoryManager{ReturnId: uuid.MustParse(id)}
	return Handler{GameManager: &gameManager}
}

func fakeContext(method , path, payload string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(payload))
	recorder := httptest.NewRecorder()
	return e.NewContext(req, recorder), recorder
}

func getPayloadFromResponse(recorder *httptest.ResponseRecorder) Response {
	decoder := json.NewDecoder(recorder.Body)
	var res Response
	decoder.Decode(&res)
	return res
}


