package quoridor

import (
	"boardbots/util"
	"github.com/google/uuid"
	"errors"
	"time"
)
type PlayerPosition int
const (
	PlayerOne   PlayerPosition = iota
	PlayerTwo
	PlayerThree
	PlayerFour
)
const BoardSize int = 17
var WinningPositions = map[PlayerPosition]util.Position{
	PlayerOne:   {Row: 16, Col: -1}, //PLAYER_ONE
	PlayerTwo:   {Col: -1},          //PLAYER_TWO
	PlayerThree: {Row: -1, Col: 16},
	PlayerFour:  {Row: -1},
}

var directions = []util.Position{
	{Row: 1},
	{Col: 1},
	{Row: -1},
	{Col: -1},
}

type Board map[util.Position]*Piece

type Player struct {
	Barriers int
	Pawn *Piece
	PlayerId uuid.UUID
	PlayerName string
}

type Game struct {
	Board Board
	Players map[PlayerPosition]*Player
	Id uuid.UUID
	CurrentTurn PlayerPosition
	StartDate, EndDate time.Time
	Winner PlayerPosition
}

type Piece struct {
	Position util.Position
	Owner PlayerPosition
}

func (board Board) getValidMoveByDirection(pawn, direction util.Position) []util.Position {
	//check if barrier in direction
	cursor := util.Position{Row: pawn.Row + direction.Row, Col: pawn.Col + direction.Col}
	if _, barrierPresent := board[cursor] ; barrierPresent {
		return nil
	}
	cursor.Row = cursor.Row + direction.Row
	cursor.Col = cursor.Col + direction.Col
	// check for pawn
	validPositions := make([]util.Position, 0, 2)
	if _, pawnPresent := board[cursor] ; pawnPresent {
		// check for possible jumps
		if _, barrierBeyondPawn := board[util.Position{ Row: cursor.Row + direction.Row, Col: cursor.Col + direction.Col}];
			barrierBeyondPawn {
			// look at diagonals instead
			validPositions = append(validPositions, getDiagonalPositions(direction, cursor, board)...)
		} else { // no barrier, final check for a pawn.
			jumpPos := util.Position{Row: cursor.Row + 2*direction.Row, Col: cursor.Col + (2 * direction.Col)}
			_, finalPawn := board[jumpPos]
			if !finalPawn && isOnBoard(jumpPos) {
				validPositions = append(validPositions, jumpPos)
			}
		}
	} else if isOnBoard(cursor){
		validPositions = append(validPositions, cursor)
	}
	return validPositions
}

func getDiagonalPositions(vector util.Position, cursor util.Position, board Board) []util.Position {
	validPositions := make([]util.Position, 0, 2)
	leftVector := util.Position{Row: -1 * vector.Col, Col: -1 * vector.Row}
	leftTurn := getTurnPosition(leftVector, cursor, board)
	if leftTurn.Row != -1 {
		validPositions = append(validPositions, leftTurn)
	}
	rightVector := util.Position{Row: vector.Col, Col: vector.Row}
	rightTurn := getTurnPosition(rightVector, cursor, board)
	if rightTurn.Row != -1 {
		validPositions = append(validPositions, rightTurn)
	}
	return validPositions
}

func getTurnPosition(vector util.Position, cursor util.Position, board Board) util.Position {
	turnCursor := util.Position{Row: cursor.Row + vector.Row, Col: cursor.Col + vector.Col}
	_, turnBarrier := board[turnCursor]
	if !turnBarrier {
		turnCursor.Row += vector.Row
		turnCursor.Col += vector.Col

		if _, turnPawn := board[turnCursor]; !turnPawn && isOnBoard(turnCursor) {
			return turnCursor
		}
	}
	return util.Position{Row: -1, Col : -1}
}
func isOnBoard(position util.Position) bool {
	return !(position.Row < 0 || position.Row >= BoardSize || position.Col < 0 || position.Col >= BoardSize)
}

func (game *Game) GetPlayer(playerId uuid.UUID) *Player {
	for _, player :=  range game.Players {
		if player.PlayerId == playerId {
			return player
		}
	}
	return nil
}
func (game *Game) AddPlayer(playerId uuid.UUID) (PlayerPosition, error) {
	for playerNumber := PlayerOne; playerNumber <= PlayerFour; playerNumber++ {
		if game.Players[playerNumber].PlayerId == uuid.Nil {
			game.Players[playerNumber].PlayerId = playerId
			return playerNumber, nil
		}
		// TODO(rwsargent) Play against yourself?
	}
	return -1, errors.New("no open player positions in this game")
}
func (game *Game) AddPiece(piece *Piece, position util.Position) {
	game.Board[position] = piece
}
func (game *Game) NextTurn() {
	next := int(game.CurrentTurn +1) % len(game.Players)
	game.CurrentTurn = PlayerPosition(next)
}

func (game *Game) IsGameOver() bool {
	return !game.EndDate.IsZero()
}

func NewPiece(position util.Position) *Piece {
	piece := new(Piece)
	piece.Position = position
	return piece
}

func (game *Game) MaybeReturnWinnerPlayerPosition() PlayerPosition {
	for position, player := range game.Players {
		winningPosition := WinningPositions[position]
		if player.Pawn.Position.Row == winningPosition.Row ||
			player.Pawn.Position.Col == winningPosition.Col {
			return position
		}
	}
	return -1
}

