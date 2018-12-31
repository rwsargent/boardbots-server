package transport

import (
	"boardbots/quoridor"
	"boardbots/util"
	"sort"
	"net/http"
	"github.com/labstack/echo"
	"github.com/google/uuid"
)

type (
	BaseResponse struct {
		Error string `json:"error"`
	}

	BoardResponse struct {
		Board []TPiece `json:"board"`
		CurrentTurn quoridor.PlayerPosition `json:"currentTurn"`
	}

	GameResponse struct {
		BaseResponse
		BoardResponse
		Players []TPlayer
	}

	GameRequest struct {
		GameId uuid.UUID `json:"gameId"`
	}

	TPiece struct {
		Type rune `json:"type"`
		Position util.Position `json:"position"`
		Owner quoridor.PlayerPosition `json:"owner"`
	}

	TPlayer struct {
		Barriers int `json:"barriers"`
		PlayerName string `json:"playerName"`
		PawnPosition util.Position `json:"pawnPosition"`
	}

	BoardOrder []TPiece
)

func StandardBadRequestError(err error) error {
	return echo.NewHTTPError(http.StatusBadRequest, BaseResponse{
		Error: err.Error(),
	})
}

func GameToGameResponse(game *quoridor.Game) GameResponse {
	return GameResponse{
	  BoardResponse: BoardResponse{
			Board:  BoardToTPieces(game.Board),
			CurrentTurn: game.CurrentTurn,
		},
		Players: toTPlayers(game.Players),
	}
}

func toTPlayers(players map[quoridor.PlayerPosition]*quoridor.Player) []TPlayer {
	tPlayers := make([]TPlayer, len(players))
	for playerPosition := int(quoridor.PlayerOne); playerPosition < len(players); playerPosition++ {
		player := players[quoridor.PlayerPosition(playerPosition)]
		tPlayers[playerPosition].Barriers = player.Barriers
		tPlayers[playerPosition].PlayerName = player.PlayerName
		tPlayers[playerPosition].PawnPosition = player.Pawn.Position
	}
	return tPlayers
}

func BoardToTPieces(board quoridor.Board) []TPiece {
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

func getPieceType(piece *quoridor.Piece) rune {
	var pieceType rune
	if quoridor.IsValidPawnLocation(piece.Position) {
		pieceType = 'p'
	} else {
		pieceType = 'b'
	}
	return pieceType
}

func (a BoardOrder) Len() int {return len(a)}
func (a BoardOrder) Swap(i, j int) {a[i], a[j] = a[j], a[i] }
func (a BoardOrder) Less(left, right int) bool {
	if a[left].Position.Row == a[right].Position.Row {
		return a[left].Position.Col < a[right].Position.Col
	}
	return a[left].Position.Row < a[right].Position.Row
}
