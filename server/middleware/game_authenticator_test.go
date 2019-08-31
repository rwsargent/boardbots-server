package middleware

import (
	"boardbots-server/manager"
	"boardbots-server/quoridor"
	"boardbots-server/server/context"
	tu "boardbots-server/server/testingutils"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var Authenticator = GameAuthenticator{
	GameManager: &manager.FakeMemoryManager{},
}
var passedContext echo.Context
var next = func(ctx echo.Context) error {
	passedContext = ctx
	return nil
}

func Test_GameAuthenticator(t *testing.T) {
	ctx, _ := tu.FakeContext(http.MethodPost, "/", `{"gameId":"`+tu.TestUUID.String()+"\"}")

	bbCtx := context.DefaultBBContext{
		Context: ctx,
		PlayerPrinciple: context.PlayerPrinciple{
			UserId:   tu.TestUUID,
			UserName: "Ryan!",
			Password: "Password",
		},
		Game: nil,
	}

	authHandler := GameAuthenticator{
		&manager.FakeMemoryManager{
			ReturnGame: quoridor.NewTwoPersonGame(),
		},
	}

	handler := authHandler.Authenticate(next)
	result := handler(bbCtx)

	assert.NoError(t, result)
}
