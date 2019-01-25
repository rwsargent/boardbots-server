package newuser

import (
	"boardbots/server/persistence"
	"github.com/labstack/echo"
	"net/http"
	"boardbots/server/transport"
)

type (
	Handler struct {
		UserPortal persistence.UserPortal
	}

	Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func (h Handler) NewUser(ctx echo.Context) error {
	req := new(Request)
	if err := ctx.Bind(req); err != nil {
		return transport.StandardBadRequestError(err)
	}
	if err := h.UserPortal.NewUser(req.Username, req.Password); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func ApplyRoute(server *echo.Echo, portal persistence.UserPortal) {
	server.POST("/newuser", Handler{portal}.NewUser)
}