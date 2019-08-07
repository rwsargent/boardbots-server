package makegame

import (
	"boardbots-server/manager"
	"boardbots-server/quoridor"
	"boardbots-server/server/context"
	"boardbots-server/server/transport"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

type (
	Handler struct {
		GameManager manager.GameManager
	}
	Response struct {
		transport.BaseResponse
		GameId uuid.UUID `json:"gameId"`
	}
)

func ApplyRoute(server *echo.Group, gameManager manager.GameManager) {
	h := Handler{
		GameManager: gameManager,
	}

	server.POST("/makegame", h.MakeGame)
}

func (h *Handler) MakeGame(ctx echo.Context) error {
	bbCtx := ctx.(*context.DefaultBBContext)
	playerName := bbCtx.GetPlayerName()
	response := Response{}
	if len(playerName) == 0 {
		response.Error = "expected a player name, didn't find one"
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	game := quoridor.NewTwoPersonGame(uuid.New())
	playerPosition, err := game.AddPlayer(bbCtx.PlayerPrinciple.UserId, bbCtx.PlayerPrinciple.UserName)
	if err != nil {
		return transport.StandardBadRequestError(err)
	}
	if playerPosition != quoridor.PlayerOne {
		log.Errorf("new game can't add player")
		return transport.HandledServerError(errors.New(fmt.Sprintf("new game can't add player, expecting 0 got %d", playerPosition)))
	}
	err = h.GameManager.AddGame(*game)
	if err != nil {
		response.Error = "error saving game, " + err.Error()
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response.GameId = game.Id
	return ctx.JSON(http.StatusOK, response)
}
