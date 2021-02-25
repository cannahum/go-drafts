package player

import (
	"github.com/cannahum/go-drafts/003-TicTacToe/gameboard"
	"github.com/cannahum/go-drafts/003-TicTacToe/move"
)

// Player is an interface that makes moves on the game board.
type Player interface {
	// makeAMove(move)
	makeAMove(move.Coordinates) (*gameboard.MoveInGame, error)
}
