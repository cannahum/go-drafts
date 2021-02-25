package gameboard

import (
	"errors"

	"github.com/cannahum/go-drafts/003-TicTacToe/move"
)

type moveInGame struct {
	moveType    move.TicTacToeMoveType
	coordinates move.Coordinates
}

type gameboard struct {
	moves []moveInGame
	// board
}

func (gb *gameboard) getBoardStateAt(moveNumber int) {

}

func (gb *gameboard) getCurrentBoardState() {

}

func (gb *gameboard) registerMove(moveType move.TicTacToeMoveType, coor move.Coordinates) error {
	currentListOfMoves := gb.moves

	if !isValidMove(currentListOfMoves, moveType, coor) {
		return errors.New("Invalid move")
	}

	currentListOfMoves = append(currentListOfMoves, moveInGame{moveType, coor})
	gb.moves = currentListOfMoves

	return nil
}

func NewGameBoard() GameBoard {
	return &gameboard{[]moveInGame{}}
}

func isValidMove(listOfMoves []moveInGame, moveType move.TicTacToeMoveType, coor move.Coordinates) bool {
	return true
}
