package gameboard

import "github.com/cannahum/go-drafts/003-TicTacToe/move"

// GameBoard interface is a way to interact with an existing board instance
type GameBoard interface {
	registerMove(move.TicTacToeMoveType, move.Coordinates) error
	getCurrentBoardState()
	getBoardStateAt(int)
}
