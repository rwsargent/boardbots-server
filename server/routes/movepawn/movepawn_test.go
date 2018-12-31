package movepawn

import (
	"testing"

	"net/http"
	"boardbots/server/context"
	"boardbots/quoridor"
	tu "boardbots/server/testingutils"
	"github.com/google/uuid"
	"fmt"
)

func TestHandler_MovePawn(t *testing.T) {
	payload := fmt.Sprintf(`{"gameId":%s,"position" : {"row" : 2, "col" : 16"}}`, tu.TestUUID.String())
	ctx, _ := tu.FakeContext(http.MethodPost,"/movepawn", payload)
	bbCtx := context.DefaultBBContext{
		ctx,
		context.PlayerPrinciple{UserName:"name", UserId:tu.TestUUID}}
	game := quoridor.NewTwoPersonGame()
	game.AddPlayer(bbCtx.PlayerPrinciple.UserId)

	handler := fakeHandler(tu.TestUUID, game)

	handler.MovePawn(ctx)


}


func fakeHandler(id uuid.UUID, game *quoridor.Game) Handler {
	return Handler{
	}
}