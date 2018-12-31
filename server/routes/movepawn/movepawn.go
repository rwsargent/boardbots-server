package movepawn

import (
	"boardbots/manager"
	"boardbots/server/transport"
	"github.com/google/uuid"
	"boardbots/util"
	"github.com/labstack/echo"
	"boardbots/server/context"
)

type (
	Handler struct {
		GameManager manager.GameManager
	}
	Request struct {
		GameId uuid.UUID
		Position util.Position
	}
	GameResponse struct {
		transport.GameResponse
	}
)

func (h Handler) MovePawn(ctx echo.Context) error {
	bbCtx := ctx.(context.DefaultBBContext)
	req := getRequest(bbCtx);
	game := h.GameManager.GetGame()

	return bbCtx.JSON()
}

func getRequest(bbContext context.DefaultBBContext) Request {

}