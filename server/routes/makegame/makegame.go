package makegame

import (
	"github.com/labstack/echo"
	"net/http"
	"boardbots/manager"
	"boardbots/quoridor"
	"github.com/google/uuid"
	"boardbots/server/context"
	"boardbots/server/transport"
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
	bbCtx := ctx.(context.DefaultBBContext)
	playerName := bbCtx.GetPlayerName()
	response := Response{}
	if len(playerName) == 0 {
		response.Error = "expected a player name, didn't find one"
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	game := quoridor.NewTwoPersonGame()
	game.Players[quoridor.PlayerOne].PlayerName = playerName
	gameId, err  := h.GameManager.AddGame(game)
	if err != nil {
		response.Error = "error saving game, " + err.Error()
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response.GameId = gameId
	return ctx.JSON(http.StatusOK, response)
}