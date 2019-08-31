package movepawn

import (
	"boardbots-server/manager"
	"boardbots-server/server/context"
	"boardbots-server/server/transport"
	"boardbots-server/util"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type (
	Handler struct {
		GameManager manager.GameManager
	}
	Request struct {
		GameId   uuid.UUID
		Position util.Position
	}
	GameResponse struct {
		transport.GameResponse
	}
)

func (h Handler) MovePawn(ctx echo.Context) error {
	bbCtx := ctx.(context.DefaultBBContext)
	req := getRequest(bbCtx)
	game := h.GameManager.GetGame()

	return bbCtx.JSON()
}

func getRequest(bbContext context.DefaultBBContext) Request {

}
