package player

import (
	"github.com/cannahum/go-drafts/003-TicTacToe/gameboard"
	"github.com/cannahum/go-drafts/003-TicTacToe/move"
)

type player struct {
	moveType move.TicTacToeMoveType
}

func (p *player) MakeAMove(row, col int) *gameboard.MoveInGame {
	return &gameboard.MoveInGame{
		MoveType: p.moveType,
		Coordinates: move.Coordinates{
			Row: row,
			Col: col,
		},
	}
}

// NewPlayer is a factory method that returns a pointer to a new instance of player which implements Player interface
func NewPlayer(t move.TicTacToeMoveType) Player {
	return &player{t}
}
