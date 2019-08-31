package getgames

import (
	"boardbots-server/manager"
	"boardbots-server/server/transport"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

const GetGamesRoute = "/getgames"

type (
	Request struct {
		Status  string
		Players []string
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
	var req Request
	if err := ctx.Bind(&req); err != nil {
		log.Error(err)
		return transport.StandardBadRequestError(err)
	}
	var games []manager.ManagedGame
	params := convertRequestToParams(req)
	games, err := h.GameManager.FindGames(params)
	if err != nil {
		return transport.HandledServerError(err)
	}

	tgames := make([]transport.TGame, 0, len(games))
	for _, game := range games {
		tgames = append(tgames, transport.NewTGame(game.Game))
	}
	resp := Response{
		Games: tgames,
	}
	return ctx.JSON(200, resp)
}

func convertRequestToParams(request Request) manager.GameParameters {
	return manager.GameParameters{
		Players:   request.Players,
		GameState: request.Status,
	}
}

func ApplyRoute(group *echo.Group, gameManager manager.GameManager) {
	h := Handler{
		GameManager: gameManager,
	}

	group.POST(GetGamesRoute, h.Handle)
}
