package gameboard

import "github.com/cannahum/go-drafts/003-TicTacToe/move"

// GameBoard interface is a way to interact with an existing board instance
type GameBoard interface {
	RegisterMove(move.TicTacToeMoveType, move.Coordinates) (bool, error)
	GetCurrentBoardState() BoardRepresentation
	GetBoardStateAt(int) BoardRepresentation
}
