package transport

import (
	"boardbots-server/quoridor"
	"boardbots-server/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

const BarrierSize = 3

func Test_GameTranslates(t *testing.T) {
	game := quoridor.NewTwoPersonGame()
	tgame := NewTGame(*game)

	assert.Len(t, tgame.Board, 2)
}

func Test_GameWithBarriers(t *testing.T) {
	game := quoridor.NewTwoPersonGame()
	game.PlaceBarrier(util.Position{Row: 1, Col: 2}, quoridor.PlayerOne)
	game.PlaceBarrier(util.Position{Row: 13, Col: 10}, quoridor.PlayerTwo)
	game.PlaceBarrier(util.Position{Row: 3, Col: 8}, quoridor.PlayerOne)

	tgame := NewTGame(*game)
	addedBarriers := 3
	assert.Len(t, tgame.Board, addedBarriers*BarrierSize+2)
}

func Test_GameWithBarriers_SortedCorrectly(t *testing.T) {
	game := quoridor.NewTwoPersonGame()
	game.PlaceBarrier(util.Position{Row: 1, Col: 2}, quoridor.PlayerOne)
	game.PlaceBarrier(util.Position{Row: 13, Col: 10}, quoridor.PlayerTwo)
	game.PlaceBarrier(util.Position{Row: 3, Col: 8}, quoridor.PlayerOne)

	res := NewTGame(*game)

	prev := res.Board[0]
	for _, piece := range res.Board {
		if prev.Position.Row == piece.Position.Row {
			assert.Truef(t, prev.Position.Col <= piece.Position.Col,
				"%v out of order with %v", prev.Position, piece.Position)
		} else {
			assert.True(t, prev.Position.Row < piece.Position.Row,
				"%v out of order with %v", prev.Position, piece.Position)
		}
		prev = piece
	}
}

func Test_TwoGame_PlayersCorrect(t *testing.T) {
	game := quoridor.NewTwoPersonGame()
	game.Players[quoridor.PlayerOne].PlayerName = "foo"
	game.Players[quoridor.PlayerTwo].PlayerName = "bar"

	res := NewTGame(*game)

	assert.Len(t, res.Players, 2)
	exp := struct {
		n string
		p util.Position
	}{"foo", util.Position{0, 8}}
	checkPlayer(t, res.Players[0], exp)
	exp = struct {
		n string
		p util.Position
	}{"bar", util.Position{16, 8}}
	checkPlayer(t, res.Players[1], exp)
}

func checkPlayer(t *testing.T, player TPlayerState, expected struct {
	n string
	p util.Position
}) {
	assert.Equal(t, player.PlayerName, expected.n)
	assert.Equal(t, player.PawnPosition, expected.p)
}
