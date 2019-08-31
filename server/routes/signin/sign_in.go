package signin

import (
	"boardbots-server/server/persistence"
	"boardbots-server/server/transport"
	"github.com/labstack/echo"
	"net/http"
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

func (h Handler) SignIn(ctx echo.Context) error {
	req := new(Request)
	if err := ctx.Bind(req); err != nil {
		return transport.StandardBadRequestError(err)
	}
	if ok := h.UserPortal.ValidateCredentials(req.Username, req.Password); ok {
		return ctx.NoContent(http.StatusOK)
	} else {
		return transport.NewHandledError(http.StatusForbidden, "invalid username or password")
	}
}

func ApplyRoute(server *echo.Echo, portal persistence.UserPortal) {
	server.POST("/signin", Handler{portal}.SignIn)
}
