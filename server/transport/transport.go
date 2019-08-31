package transport

import (
	quor "boardbots-server/quoridor"
	"boardbots-server/util"
	"github.com/google/uuid"
	"net/http"
	"sort"
)

type (
	Transportable interface {
		ToTee() interface{}
	}

	BaseCommand struct {
		GameId   uuid.UUID     `json: "gameId"`
		Position util.Position `json: "position"`
	}

	TPlayerState struct {
		PlayerName   string        `json:"playerName"`
		PawnPosition util.Position `json:"pawnPosition"`
		Barriers     int           `json:"barriers"`
	}

	TGame struct {
		GameId      uuid.UUID      `json: "gameId"`
		Players     []TPlayerState `json: "players"`
		CurrentTurn int            `json:"currentTurn'"`
		StartDate   int64          `json:"startDate"`
		EndDate     int64          `json:"startDate"`
		Winner      int            `json:"winner"`
		Board       TBoard
	}

	TBoard     []TPiece
	BoardOrder []TPiece // For sorting readability.

	BaseResponse struct {
		Error string `json:"error"`
	}

	BoardResponse struct {
		Board       []TPiece            `json:"board"`
		CurrentTurn quor.PlayerPosition `json:"currentTurn"`
	}

	GameResponse struct {
		BaseResponse
		Game TGame `json:"game"`
	}

	GameRequest struct {
		GameId uuid.UUID
	}

	TPiece struct {
		Type     rune                `json:"type"`
		Position util.Position       `json:"position"`
		Owner    quor.PlayerPosition `json:"owner"`
	}
)

func StandardBadRequestError(err error) error {
	return HandledError{http.StatusBadRequest, err.Error()}
}

func HandledServerError(err error) error {
	return HandledError{http.StatusInternalServerError, err.Error()}
}

func ToTPlayers(players map[quor.PlayerPosition]*quor.Player) []TPlayerState {
	tPlayers := make([]TPlayerState, len(players))
	for playerPosition := int(quor.PlayerOne); playerPosition < len(players); playerPosition++ {
		player := players[quor.PlayerPosition(playerPosition)]
		tPlayers[playerPosition].Barriers = player.Barriers
		tPlayers[playerPosition].PlayerName = player.PlayerName
		tPlayers[playerPosition].PawnPosition = player.Pawn.Position
	}
	return tPlayers
}

func BoardToTBoard(board quor.Board) TBoard {
	pieces := make([]TPiece, 0, len(board))
	for position, piece := range board {
		pieceType := getPieceType(piece)
		pieces = append(pieces, TPiece{
			Type:     pieceType,
			Position: position,
			Owner:    piece.Owner,
		})
	}
	sort.Sort(BoardOrder(pieces))
	return pieces
}

func getPieceType(piece quor.Piece) rune {
	var pieceType rune
	if quor.IsValidPawnLocation(piece.Position) {
		pieceType = 'p'
	} else {
		pieceType = 'b'
	}
	return pieceType
}

func NewTGame(game quor.Game) TGame {
	tgame := TGame{
		GameId:    game.Id,
		Players:   NewTPlayerStates(game),
		EndDate:   game.EndDate.Unix(),
		StartDate: game.StartDate.Unix(),
		Board:     BoardToTBoard(game.Board),
	}
	return tgame
}

func NewTPlayerStates(game quor.Game) []TPlayerState {
	playerStates := make([]TPlayerState, 0, len(game.Players))
	for playerNum := int(quor.PlayerOne); playerNum < len(game.Players); playerNum++ {
		player := game.Players[quor.PlayerPosition(playerNum)]
		playerState := TPlayerState{
			PlayerName:   player.PlayerName,
			PawnPosition: player.Pawn.Position,
		}
		playerStates = append(playerStates, playerState)
	}
	return playerStates
}

func (a BoardOrder) Len() int      { return len(a) }
func (a BoardOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BoardOrder) Less(left, right int) bool {
	if a[left].Position.Row == a[right].Position.Row {
		return a[left].Position.Col < a[right].Position.Col
	}
	return a[left].Position.Row < a[right].Position.Row
}
