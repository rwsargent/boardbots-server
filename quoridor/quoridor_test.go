package quoridor

import (
	tu "boardbots/server/testingutils"
	u "boardbots/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoPlayerGame(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)

	assertBaseGame(t, game)
	assertCorrectPlayerInit(t, game)
	assertNoExtraPlayersCreated(t, game)
}

func Test_FirstMoveIsValid(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	board, err := game.MovePawn(u.Position{2, 8}, PlayerOne)
	assert.Nil(t, err, "Valid Move")
	assert.Len(t, board, 2, "Game has two pawns")

	assert.Equal(t, game.Players[PlayerOne].Pawn.Position, u.Position{2, 8})
}

func Test_FirstMoveIsAnInvalidMove(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	_, err := game.MovePawn(u.Position{4, 4}, PlayerOne)
	assert.EqualError(t, err, "the pawn cannot reach that square")
}

func Test_PlaceBarrier(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	placePosition := u.Position{1, 6}
	board, err := game.PlaceBarrier(placePosition, PlayerOne)
	assert.NotNil(t, board[placePosition])
	assert.NotNil(t, board[u.Position{1, 7}])
	assert.NotNil(t, board[u.Position{1, 8}])
	_, present := board[u.Position{1, 9}]
	assert.False(t, present)
	assert.Nil(t, err, "Valid placement")
}

func Test_PlaceBarrierWithNoMoreBarriers(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	for i := 0; i < 18; i++ {
		row := 14 - (4 * (i / 8)) // offset after a row as filled up
		col := 1 + (2 * (i % 8))  // wrap around to the same col, since row changes.
		position := u.Position{Row: row, Col: col}
		board, err := game.PlaceBarrier(position, PlayerPosition(i%2))
		assert.Nil(t, err, "No error expected")
		assert.NotNil(t, board[position], "board should be placed")
		assert.NotNil(t, game.Board[position], "board should be placed")
	}
	game.PlaceBarrier(u.Position{Row: 0, Col: 3}, PlayerOne)
	board, err := game.PlaceBarrier(u.Position{Row: 0, Col: 1}, PlayerTwo)
	assert.Nil(t, err, "No error expected")
	assert.NotNil(t, board[u.Position{Row: 0, Col: 1}], "board should be placed")

	illegalPosition := u.Position{Row: 0, Col: 7}
	board, err = game.PlaceBarrier(illegalPosition, PlayerOne)
	assert.NotNil(t, err, "expect 11th barrier to throw error")
	_, ok := board[illegalPosition]
	assert.False(t, ok, "barrier should not be on board")
	_, ok = game.Board[illegalPosition]
	assert.False(t, ok)
}

func Test_PawnCannotMoveOverBarrier(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	game.PlaceBarrier(u.Position{1, 8}, PlayerOne)
	game.PlaceBarrier(u.Position{5, 6}, PlayerTwo)
	_, err := game.MovePawn(u.Position{2, 8}, PlayerOne)
	assert.NotNil(t, err, "Expected error")
	assert.Equal(t, "the pawn cannot reach that square", err.Error(), "Wrong message")
}

func Test_OneValidDirection(t *testing.T) {
	board := make(Board)
	placePawn(board)
	placeBarrier(board)

	moves := board.GetValidPawnMoves(u.Position{0, 0})
	assert.Len(t, moves, 1)
	assert.Equal(t, moves[0], u.Position{2, 0})
}

func Test_JumpPawn(t *testing.T) {
	board := make(Board)
	setupPawnBlockingBoard(board)
	moves := board.GetValidPawnMoves(u.Position{0, 8})
	assert.Len(t, moves, 1)
	assert.Equal(t, moves[0], u.Position{4, 8})
}

func Test_JumpDiagonal(t *testing.T) {
	board := make(Board)
	board[u.Position{2, 8}] = Piece{}
	board[u.Position{4, 8}] = Piece{}

	//Barriers
	board[u.Position{0, 9}] = Piece{}
	board[u.Position{1, 9}] = Piece{}
	board[u.Position{2, 9}] = Piece{}

	board[u.Position{0, 7}] = Piece{}
	board[u.Position{1, 7}] = Piece{}
	board[u.Position{2, 7}] = Piece{}

	board[u.Position{5, 8}] = Piece{}
	board[u.Position{5, 9}] = Piece{}
	board[u.Position{5, 10}] = Piece{}

	moves := board.GetValidPawnMoves(u.Position{2, 8})
	assert.Len(t, moves, 3)

	expecteds := map[u.Position]bool{
		{4, 10}: true,
		{4, 6}:  true,
		{0, 8}:  true,
	}
	for _, move := range moves {
		if _, ok := expecteds[move]; !ok {
			t.Errorf("Missing expected move %v", move)
		}
	}
}

func Test_FourDirectionsForFree(t *testing.T) {
	board := make(Board)
	board[u.Position{2, 8}] = Piece{}

	moves := board.GetValidPawnMoves(u.Position{2, 8})

	assert.Len(t, moves, 4)
}

func Test_GameNotOver(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	assert.Equal(t, PlayerPosition(-1), game.MaybeReturnWinnerPlayerPosition())
}

func Test_GameOver(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	p1Pawn := &game.Players[PlayerOne].Pawn
	delete(game.Board, p1Pawn.Position)
	p1Pawn.Position = u.Position{16, 8}
	game.Board[p1Pawn.Position] = *p1Pawn

	assert.Equal(t, PlayerOne, game.MaybeReturnWinnerPlayerPosition())
}

func Test_BarrierPlacementBlocksWin(t *testing.T) {
	var board = `.......|0|.......
.......|.|.......
.......|.|.......
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
.................
................`

	game, err := BuildQuoidorBoardFromString(board)
	if err != nil {
		t.Error(err.Error())
	}

	newBoard, err := game.PlaceBarrier(u.Position{3, 8}, PlayerOne)
	assert.NotNil(t, err)
	assertNoPiece(t, newBoard, u.Position{Row: 3, Col: 8})
	assertNoPiece(t, newBoard, u.Position{Row: 3, Col: 9})
	assertNoPiece(t, newBoard, u.Position{Row: 3, Col: 10})
}

func assertNoPiece(t *testing.T, board Board, position u.Position) {
	if _, exists := board[position]; exists {
		t.Error("found unexpected piece at Position: ", position)
	}
}

func TestAddPlayer(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	playerId := uuid.New()
	player, err := game.AddPlayer(playerId, "Test Name")
	assert.NoError(t, err)
	assert.Equal(t, PlayerOne, player)

	player, err = game.AddPlayer(uuid.New(), "Test Name")
	assert.NoError(t, err)
	assert.Equal(t, PlayerTwo, player)
}

func Test_NewGame_PlayerOnesTurn(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	assert.Equal(t, game.CurrentTurn, PlayerOne)
}

func Test_MovePawn_FailsIfWrongPlayer(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	_, err := game.MovePawn(u.Position{14, 8}, PlayerTwo)
	assert.Error(t, err)
}

func Test_MovePawn_TurnCheck(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	game.MovePawn(u.Position{2, 8}, PlayerOne)

	assert.Equal(t, game.CurrentTurn, PlayerTwo)
}

func Test_MovePawnTwice_TurnCheck(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)
	game.MovePawn(u.Position{2, 8}, PlayerOne)
	game.MovePawn(u.Position{14, 8}, PlayerTwo)

	assert.Equal(t, game.CurrentTurn, PlayerOne)
}

func Test_MovePawnFourTimes_FourPersonGame(t *testing.T) {
	var err error
	game := NewFourPersonGame(tu.TestUUID)
	game.MovePawn(u.Position{2, 8}, PlayerOne)
	game.MovePawn(u.Position{14, 8}, PlayerTwo)
	_, err = game.MovePawn(u.Position{8, 2}, PlayerThree)
	assert.NoError(t, err)
	_, err = game.MovePawn(u.Position{8, 14}, PlayerFour)
	assert.NoError(t, err)

	assert.Equal(t, game.CurrentTurn, PlayerOne)
}

func Test_PlaceBarrier_WrongTurn(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)

	_, err := game.PlaceBarrier(u.Position{1, 2}, PlayerTwo)

	assert.Error(t, err)
	assert.EqualError(t, err, "wrong turn, current turn is for Player: 0")
}

func Test_PlaceBarrierChangesTurnTwoPlayer(t *testing.T) {
	game := NewTwoPersonGame(tu.TestUUID)

	game.PlaceBarrier(u.Position{1, 2}, PlayerOne)

	assert.Equal(t, game.CurrentTurn, PlayerTwo)
}

func Test_MultiplePlaceBarrier(t *testing.T) {
	game := NewFourPersonGame(tu.TestUUID)

	game.PlaceBarrier(u.Position{1, 2}, PlayerOne)
	game.PlaceBarrier(u.Position{3, 2}, PlayerTwo)
	game.PlaceBarrier(u.Position{5, 2}, PlayerThree)
	game.PlaceBarrier(u.Position{7, 2}, PlayerFour)

	assert.Equal(t, game.CurrentTurn, PlayerOne)

}

func setupPawnBlockingBoard(board Board) {
	//Pawns
	board[u.Position{0, 8}] = Piece{}
	board[u.Position{2, 8}] = Piece{}

	//Barriers
	board[u.Position{0, 9}] = Piece{}
	board[u.Position{1, 9}] = Piece{}
	board[u.Position{2, 9}] = Piece{}
	board[u.Position{0, 7}] = Piece{}
	board[u.Position{1, 7}] = Piece{}
	board[u.Position{2, 7}] = Piece{}
}

func placeBarrier(board Board) {
	board[u.Position{0, 1}] = Piece{}
	board[u.Position{1, 1}] = Piece{}
	board[u.Position{1, 1}] = Piece{}
}

func placePawn(board Board) {
	board[u.Position{0, 0}] = Piece{}
}

func assertNoExtraPlayersCreated(t *testing.T, game *Game) {
	assert.Nil(t, game.Players[PlayerThree])
	assert.Nil(t, game.Players[PlayerFour])
}

func assertCorrectPlayerInit(t *testing.T, game *Game) {
	assert.NotNil(t, game.Players[PlayerOne], "Initialized Player one!")
	assert.NotNil(t, game.Players[PlayerTwo], "Initialized Player two!")
	assert.Equal(t, u.Position{0, 8}, game.Players[PlayerOne].Pawn.Position)
	assert.Equal(t, u.Position{16, 8}, game.Players[PlayerTwo].Pawn.Position)
	assert.Equal(t, 10, game.Players[PlayerOne].Barriers)
	assert.Equal(t, 10, game.Players[PlayerTwo].Barriers)
}

func assertBaseGame(t *testing.T, game *Game) {
	assert.NotNil(t, game.Board, "Board should be initialized")
	assert.NotNil(t, game.Board, "Pieces of board should be initialized")
	assert.Equal(t, 2, len(game.Players), "Should be 4 players")
}
