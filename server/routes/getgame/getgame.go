package getgame

import (
	"boardbots/manager"
	"boardbots/server/context"
	"boardbots/server/transport"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

type (
	Handler struct {
		GameManager manager.GameManager
	}
	Request struct {
		GameId uuid.UUID
	}
	GameResponse struct {
		transport.GameResponse
	}
)

func ApplyRoute(group *echo.Group, gameManager manager.GameManager) {
	h := Handler{
		GameManager: gameManager,
	}
	group.POST("/getgame", h.GetGame)
}

func (h Handler) GetGame(ctx echo.Context) error {
	bbCtx := ctx.(*context.DefaultBBContext)
	gameId, err := getGameId(bbCtx)
	if err != nil {
		return transport.StandardBadRequestError(err)
	}
	game, err := h.GameManager.GetGame(gameId)
	if err != nil {
		return transport.StandardBadRequestError(err)
	}
	res := transport.NewTGame(*game)
	return ctx.JSON(http.StatusOK, res)
}

func getGameId(bbContext *context.DefaultBBContext) (uuid.UUID, error) {
	gameIdParam := bbContext.QueryParam("gameid")
	if len(gameIdParam) == 0 {
		return uuid.Nil, errors.New("no gameid query parameter")
	}
	gameId, err := uuid.Parse(gameIdParam)
	if err != nil {
		return uuid.Nil, errors.New(fmt.Sprintf("could not parse id: %s", gameIdParam))
	}
	return gameId, nil
}
