package middleware

import (
	"testing"
	"boardbots/manager"
	"boardbots/server/context"
	"github.com/labstack/echo"
	"net/http"
	"github.com/stretchr/testify/assert"
	tu "boardbots/server/testingutils"
	"boardbots/quoridor"
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
	ctx, _ := tu.FakeContext(http.MethodPost, "/", 	`{"gameId":"` + tu.TestUUID.String() + "\"}")

	bbCtx := context.DefaultBBContext{
		Context : ctx,
		PlayerPrinciple : context.PlayerPrinciple{
			UserId: tu.TestUUID,
			UserName : "Ryan!",
		},
		Game : nil,
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
