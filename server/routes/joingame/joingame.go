package joingame

import (
	"boardbots/manager"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"boardbots/server/transport"
	"fmt"
	"boardbots/quoridor"
	"boardbots/server/context"
)

type (
	Handler struct {
		GameManager manager.GameManager
	}

	Request struct {
		GameId uuid.UUID `json:"gameId"`
	}

	Response struct{
		transport.GameResponse
		PlayerPosition quoridor.PlayerPosition `json:"playerPosition"`
	}
)

func (h *Handler) JoinGame(ctx echo.Context) error {
	bbCtx := ctx.(context.DefaultBBContext)
	req := new(Request)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, transport.BaseResponse{
			Error : err.Error(),
		})
	}
	game, err := (h.GameManager).GetGame(req.GameId)
	if game == nil {
		return echo.NewHTTPError(http.StatusBadRequest, transport.BaseResponse{
			Error : fmt.Sprintf("Could not find a game with id: %s. %s", req.GameId.String(), err),
		})
	}
	player, err := game.AddPlayer(bbCtx.PlayerPrinciple.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, transport.BaseResponse{
			Error : fmt.Sprintf("Game %s as no open spots", req.GameId.String()),
		})
	}
	return ctx.JSON(http.StatusOK,  Response{
		GameResponse: transport.GameResponse{
			Game: transport.NewTGame(*game),
		},
		PlayerPosition: player,
	})
}

func ApplyRoute(group *echo.Group, gameManager manager.GameManager) {
	h := Handler{
		GameManager: gameManager,
	}

	group.POST("/joingame", h.JoinGame)
}