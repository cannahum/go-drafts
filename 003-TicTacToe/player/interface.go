package player

import (
	"github.com/cannahum/go-drafts/003-TicTacToe/gameboard"
)

// Player is an interface that makes moves on the game board.
type Player interface {
	// makeAMove(move)
	MakeAMove(row, col int) *gameboard.MoveInGame
}
