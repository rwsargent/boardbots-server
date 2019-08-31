package movepawn

import (
	"testing"

	"boardbots-server/quoridor"
	"boardbots-server/server/context"
	tu "boardbots-server/server/testingutils"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func TestHandler_MovePawn(t *testing.T) {
	payload := fmt.Sprintf(`{"gameId":%s,"position" : {"row" : 2, "col" : 16"}}`, tu.TestUUID.String())
	ctx, _ := tu.FakeContext(http.MethodPost, "/movepawn", payload)
	bbCtx := context.DefaultBBContext{
		ctx,
		context.PlayerPrinciple{UserName: "name", UserId: tu.TestUUID}}
	game := quoridor.NewTwoPersonGame()
	game.AddPlayer(bbCtx.PlayerPrinciple.UserId)

	handler := fakeHandler(tu.TestUUID, game)

	handler.MovePawn(ctx)

}

func fakeHandler(id uuid.UUID, game *quoridor.Game) Handler {
	return Handler{}
}
