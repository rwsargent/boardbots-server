package manager

import (
	"testing"

	q "boardbots/quoridor"
	tu "boardbots/server/testingutils"
	"github.com/stretchr/testify/assert"
)

type ManagedGamesType []ManagedGame

func TestInMemoryGameManager_AddGame(t *testing.T) {
	manager := NewMemoryGameManager()
	games := fillGames()

	for idx := 0; idx < 6; idx++ {
		assert.Nil(t, manager.AddGame(games[idx].Game), "game should add without issues")
	}

	assert.Error(t, manager.AddGame(games[0].Game))
}

func TestGetGames(t *testing.T) {
	manager := NewMemoryGameManager()
	games := fillGames()
	for idx := 0; idx < 6; idx++ {
		assert.Nil(t, manager.AddGame(games[idx].Game), "game should add without issues")
	}
	testTable := []struct {
		Name          string
		Params        GameParameters
		ExpectedGames []ManagedGame
		ExpectedError error
	}{
		{
			Name:          "Empty params, no games",
			Params:        GameParameters{},
			ExpectedGames: nil,
			ExpectedError: nil,
		},
		{
			Name: "Find all started game",
			Params: GameParameters{
				GameState: GameState_Started,
			},
			ExpectedGames: []ManagedGame{games[3]},
			ExpectedError: nil,
		},
		{
			Name: "Find games for player",
			Params: GameParameters{
				Players: []string{tu.PlayerPrinciples[0].UserId.String()},
			},
			ExpectedGames: []ManagedGame{games[0]},
			ExpectedError: nil,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			foundGames, err := manager.FindGame(tt.Params)
			if err != nil {
				if tt.ExpectedError != nil {
					assert.Equal(t, err.Error(), tt.ExpectedError.Error())
				} else {
					t.Fail()
				}
			} else {
				verifyLists(t, tt.ExpectedGames, foundGames)
			}
		})
	}
}

func fillGames() []ManagedGame {
	games := make([]ManagedGame, 0)
	for _, gameId := range tu.GameIds {
		game := q.NewTwoPersonGame(gameId)
		games = append(games, ManagedGame{
			Game:     *game,
			modCount: 0,
		})
	}

	games[0].Game.AddPlayer(tu.PlayerPrinciples[0].UserId, tu.PlayerPrinciples[0].UserName)
	games[1].Game.AddPlayer(tu.PlayerPrinciples[1].UserId, tu.PlayerPrinciples[1].UserName)
	games[2].Game.AddPlayer(tu.PlayerPrinciples[2].UserId, tu.PlayerPrinciples[2].UserName)

	// Add two players
	games[3].Game.AddPlayer(tu.PlayerPrinciples[3].UserId, tu.PlayerPrinciples[3].UserName)
	games[3].Game.AddPlayer(tu.PlayerPrinciples[4].UserId, tu.PlayerPrinciples[4].UserName)

	games[3].Game.StartGame()

	// Add two players
	games[4].Game.AddPlayer(tu.PlayerPrinciples[3].UserId, tu.PlayerPrinciples[3].UserName)
	games[4].Game.AddPlayer(tu.PlayerPrinciples[4].UserId, tu.PlayerPrinciples[4].UserName)

	return games
}
func verifyLists(t *testing.T, expected, actual []ManagedGame) {
	if len(expected) != len(actual) {
		t.Error("expected list not length of actual list")
	}

	for _, exp := range expected {
		var found bool = false
		for _, act := range actual {
			if exp.Game.Id == act.Game.Id {
				found = true
			}
		}
		if !found {
			t.Error("Could not find game", exp.Game.Id)
		}
	}
}
