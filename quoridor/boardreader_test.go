package quoridor

import (
	"boardbots/util"
	"bufio"
	"strings"
	"testing"
)

func TestBuildingBoard(t *testing.T) {
	board :=
		`........0........
.................
.................
.................
.................
.................
.................
.................
.................
.................
.................
.................
.................
......|..........
......|..........
......|.---......
........1........`

	scanner := bufio.NewReader(strings.NewReader(board))

	game, err := BuildQuoridorBoard(scanner)
	if err != nil {
		t.Errorf(err.Error())
	}
	if _, ok := game.Board[util.Position{0, 8}]; !ok {
		t.Errorf("Player one not on board created")
	}

	if _, ok := game.Board[util.Position{15, 8}]; !ok {
		t.Errorf("Missing barrier")
	}

	if _, ok := game.Board[util.Position{15, 6}]; !ok {
		t.Error("Missing barrier")
	}
	if player, ok := game.Players[PlayerOne]; ok {
		expected := util.Position{0, 8}
		if player.Pawn.Position != expected {
			t.Errorf("Pawn not created")
		}
	} else {
		t.Errorf("Player not created")
	}

	if player, ok := game.Players[PlayerTwo]; ok {
		expected := util.Position{16, 8}
		if player.Pawn.Position != expected {
			t.Errorf("Pawn not created")
		}
	} else {
		t.Errorf("Player not created")
	}

}
