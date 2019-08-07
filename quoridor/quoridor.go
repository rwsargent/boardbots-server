package quoridor

import (
	"boardbots-server/util"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func NewTwoPersonGame(id uuid.UUID) *Game {
	game := initGame(id)

	barrierCount := 10
	playerOne := newPlayer(barrierCount, util.Position{Col: 8}, PlayerOne)
	game.Players[PlayerOne] = playerOne
	game.Board[playerOne.Pawn.Position] = playerOne.Pawn

	playerTwo := newPlayer(barrierCount, util.Position{Row: 16, Col: 8}, PlayerTwo)
	game.Players[PlayerTwo] = playerTwo
	game.Board[playerTwo.Pawn.Position] = playerTwo.Pawn

	return game
}

func NewFourPersonGame(id uuid.UUID) *Game {
	game := initGame(id)
	barrierCount := 5
	game.Players[PlayerOne] = newPlayer(barrierCount, util.Position{Col: 8}, PlayerOne)
	game.Board[game.Players[PlayerOne].Pawn.Position] = game.Players[PlayerOne].Pawn

	game.Players[PlayerTwo] = newPlayer(barrierCount, util.Position{Row: 16, Col: 8}, PlayerTwo)
	game.Board[game.Players[PlayerTwo].Pawn.Position] = game.Players[PlayerTwo].Pawn

	game.Players[PlayerThree] = newPlayer(barrierCount, util.Position{Row: 8}, PlayerThree)
	game.Board[game.Players[PlayerThree].Pawn.Position] = game.Players[PlayerThree].Pawn

	game.Players[PlayerFour] = newPlayer(barrierCount, util.Position{Row: 8, Col: 16}, PlayerFour)
	game.Board[game.Players[PlayerFour].Pawn.Position] = game.Players[PlayerFour].Pawn

	return game
}

func initGame(id uuid.UUID) *Game {
	game := new(Game)
	game.Board = make(map[util.Position]Piece)
	game.Players = make(map[PlayerPosition]*Player)
	game.CurrentTurn = PlayerOne
	game.Id = id
	game.Name = id.String()
	return game
}

func newPlayer(barrierCount int, position util.Position, playerPosition PlayerPosition) *Player {
	p := new(Player)
	p.Barriers = barrierCount
	p.Pawn = Piece{Position: position, Owner: playerPosition}
	return p
}

// Return the state of the Board after the move, of the current board if an error
func (game *Game) MovePawn(newPosition util.Position, player PlayerPosition) (Board, error) {
	pawn := &game.Players[player].Pawn
	if !IsValidPawnLocation(newPosition) {
		return game.Board, errors.New("invalid Pawn Location")
	}
	if game.CurrentTurn != player {
		return game.Board, errors.New(fmt.Sprintf("wrong turn, current turn is for Player: %d", game.CurrentTurn))
	}
	if moveError := isValidPawnMove(newPosition, pawn.Position, &game.Board); moveError != nil {
		return game.Board, moveError
	}
	delete(game.Board, pawn.Position)
	pawn.Position = newPosition
	game.Board[pawn.Position] = *pawn
	checkGameOver(game)
	game.NextTurn()
	return game.Board, nil
}

func checkGameOver(game *Game) {
	if winner := game.MaybeReturnWinnerPlayerPosition(); winner != -1 {
		game.EndDate = time.Now()
		game.Winner = winner
	}
}

func (game *Game) PlaceBarrier(position util.Position, player PlayerPosition) (Board, error) {
	if game.CurrentTurn != player {
		return game.Board, errors.New(fmt.Sprintf("wrong turn, current turn is for Player: %d", game.CurrentTurn))

	}
	if invalidPosition(position) {
		return game.Board, errors.New("invalid location for a barrier")
	}
	if playerHasNoMoreBarriers(game.Players[player]) {
		return game.Board, errors.New("the player has no more barriers to play")
	}
	barrierPositions := createBarrierPositions(position)
	if barriersAreInTheWay(barrierPositions, game.Board) {
		return game.Board, errors.New("the new barrier intersects with another")
	}
	if barrierPreventsWin(barrierPositions, game) {
		return game.Board, errors.New("the barrier prevents a players victory")
	}
	game.Players[player].Barriers--
	barrier := Piece{Position: position, Owner: player}
	for _, pos := range barrierPositions {
		game.AddPiece(barrier, pos)
	}
	game.NextTurn()
	return game.Board, nil
}

func invalidPosition(position util.Position) bool {
	return position.Row&0x1 == position.Col&0x1 || // both col and row are even or odd
		// can't be on the last valid row/
		!(position.Row < BoardSize-1 &&
			(position.Col < BoardSize-1))
}

func playerHasNoMoreBarriers(player *Player) bool {
	return player.Barriers <= 0
}

func barrierPreventsWin(positions [3]util.Position, game *Game) bool {
	for _, position := range positions {
		game.Board[position] = Piece{Position: position, Owner: PlayerOne}
	}
	//remove those temporary barriers no matter what
	defer func() {
		for _, position := range positions {
			delete(game.Board, position)
		}
	}()

	for playerPosition, player := range game.Players {
		path := game.FindPath(player.Pawn.Position, WinningPositions[playerPosition])
		if path == nil {
			return true
		}
	}
	return false
}

func createBarrierPositions(position util.Position) [3]util.Position {
	var positions [3]util.Position
	if isABarrierRow(position) {
		positions = buildHorizontalBarriers(position)
	} else if isABarrierColumn(position) {
		positions = buildVerticalBarriers(position)
	}
	return positions
}

func barriersAreInTheWay(positions [3]util.Position, board Board) bool {
	for _, pos := range positions {
		if _, ok := board[pos]; ok {
			return true
		}
	}
	return false
}

func buildVerticalBarriers(position util.Position) [3]util.Position {
	return [3]util.Position{
		{position.Row + 0, position.Col},
		{position.Row + 1, position.Col},
		{position.Row + 2, position.Col},
	}
}
func buildHorizontalBarriers(position util.Position) [3]util.Position {
	return [3]util.Position{
		{position.Row, position.Col + 0},
		{position.Row, position.Col + 1},
		{position.Row, position.Col + 2},
	}
}

func isABarrierColumn(position util.Position) bool {
	return position.Col&0x1 == 1 && position.Row&0x1 == 0
}

func isABarrierRow(position util.Position) bool {
	return position.Row&0x1 == 1 && position.Col&0x1 == 0
}

func IsValidPawnLocation(position util.Position) bool {
	return position.Col%2 == 0 && position.Row%2 == 0
}

func isValidPawnMove(new util.Position, current util.Position, board *Board) error {
	validPawnMoves := board.GetValidPawnMoves(current)
	for _, validPosition := range validPawnMoves {
		if validPosition == new {
			return nil
		}
	}
	return errors.New("the pawn cannot reach that square")

	return nil
}

func (board Board) GetValidPawnMoves(pawnPosition util.Position) []util.Position {
	validPositions := make([]util.Position, 0, 6)
	// pawn goes down
	validPositions = append(validPositions,
		board.getValidMoveByDirection(pawnPosition, util.Position{1, 0})...)
	// right
	validPositions = append(validPositions,
		board.getValidMoveByDirection(pawnPosition, util.Position{0, 1})...)
	// up
	validPositions = append(validPositions,
		board.getValidMoveByDirection(pawnPosition, util.Position{-1, 0})...)
	// left
	validPositions = append(validPositions,
		board.getValidMoveByDirection(pawnPosition, util.Position{0, -1})...)
	return validPositions
}
