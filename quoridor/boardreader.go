package quoridor

import (
	"boardbots-server/util"
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func BuildQuoidorBoardFromString(board string) (*Game, error) {
	return BuildQuoridorBoard(bufio.NewReader(strings.NewReader(board)))
}

func BuildQuoridorBoard(reader *bufio.Reader) (*Game, error) {
	game := &Game{
		Board:   make(Board),
		Players: make(map[PlayerPosition]*Player),
	}
	var row = 0
	for {
		line, _, err := reader.ReadLine()
		if line == nil || err != nil {
			break
		}
		for col, char := range line {
			if char == '-' || char == '|' {
				position := util.Position{
					Row: row,
					Col: col,
				}
				game.Board[position] = Piece{Position: position}
			} else if unicode.IsDigit(rune(char)) {
				playerPos := PlayerPosition(char - '0')
				piecePos := util.Position{row, col}
				player := &Player{
					Pawn:     Piece{Position: piecePos, Owner: playerPos},
					Barriers: 10,
				}
				game.Players[playerPos] = player
				game.Board[piecePos] = player.Pawn
				col++
			} else if char != '.' {
				panic(fmt.Sprintf("Unexpected character %s", char))
			}
		}
		row++
	}
	return game, nil
}
