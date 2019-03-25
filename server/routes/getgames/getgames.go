package getgames

import (
	"boardbots/manager"
	"boardbots/quoridor"
	"boardbots/server/context"
	"boardbots/server/transport"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

const GetGamesRoute = "/getgames"

type (
	Request struct {
		Status string
	}

	Response struct {
		transport.BaseResponse
		Games []transport.TGame `json:"games"`
	}

	Handler struct {
		GameManager manager.GameManager
	}
)

func (h Handler) Handle(ctx echo.Context) error {
	bbCtx := ctx.(*context.DefaultBBContext)
	var req Request
	if err := ctx.Bind(&req); err != nil {
		log.Error(err)
		return transport.StandardBadRequestError(err)
	}
	var games []quoridor.Game
	games, err := h.GameManager.GetGamesForUser(bbCtx.GetPlayerId())
	if err != nil {
		return transport.HandledServerError(err)
	}

	tgames := make([]transport.TGame, 0, len(games))
	for _, game := range games {
		tgames = append(tgames, transport.NewTGame(game))
	}
	resp := Response{
		Games: tgames,
	}
	return ctx.JSON(200, resp)
}

func ApplyRoute(group *echo.Group, gameManager manager.GameManager) {
	h := Handler{
		GameManager: gameManager,
	}

	group.POST(GetGamesRoute, h.Handle)
}
