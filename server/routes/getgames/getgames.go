package getgames

import (
	"boardbots/server/transport"
	"boardbots/manager"
	"github.com/labstack/echo"
	"boardbots/server/context"
)
const GetGamesRoute = "/getgames"

type (
	Response struct {
		transport.BaseResponse
		Games []transport.TGame `json:"games"`
	}

	Handler struct {
		GameManager manager.GameManager
	}
)

func (h Handler)Handle(ctx echo.Context) error {
	bbCtx := ctx.(context.DefaultBBContext)
	games, err := h.GameManager.GetGamesForUser(bbCtx.GetPlayerId())
	if err != nil {
		return transport.HandledServerError(err)
	}

	tgames := make([]transport.TGame, 0, len(games))
	for _, game := range games {
		tgames = append(tgames, transport.NewTGame(game))
	}
	resp := Response{
		Games : tgames,
	}
	return ctx.JSON(200, resp)
}

func ApplyRoute(group *echo.Group, gameManager manager.GameManager) {
	h := Handler{
		GameManager: gameManager,
	}

	group.POST(GetGamesRoute, h.Handle)
}