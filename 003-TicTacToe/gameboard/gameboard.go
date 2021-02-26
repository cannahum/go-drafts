package gameboard

import (
	"github.com/cannahum/go-drafts/003-TicTacToe/move"
)

type MoveInGame struct {
	MoveType    move.TicTacToeMoveType
	Coordinates move.Coordinates
}

type gameBoard struct {
	moveRegister *moveRegister
	board        *BoardRepresentation
}

func (gb *gameBoard) GetBoardStateAt(moveNumber int) BoardRepresentation {
	// This returns a new board
	newBoard := BoardRepresentation{}
	for moveIndex, move := range gb.moveRegister.moves {
		if moveIndex == moveNumber {
			break
		}

		newBoard[move.Coordinates.Row][move.Coordinates.Col] = move.MoveType
	}

	return *&newBoard
}

func (gb *gameBoard) GetCurrentBoardState() BoardRepresentation {
	return *gb.board
}

func (gb *gameBoard) updateGameWithNewBoardState() {
	lastMoveNumber := len(gb.moveRegister.moves)
	b := gb.GetBoardStateAt(lastMoveNumber)
	gb.board = &b
}

func (gb *gameBoard) isGameOver() bool {
	boardRep := *gb.board
	return boardRep.hasColTicTacToe() ||
		boardRep.hasRowTicTacToe() ||
		boardRep.hasDiagonalTicTacToe()
}

func (gb *gameBoard) RegisterMove(newMove MoveInGame) (bool, error) {
	err := gb.moveRegister.registerMove(newMove)
	if err != nil {
		return false, err
	}
	gb.updateGameWithNewBoardState()
	return gb.isGameOver(), nil
}

func NewGameBoard() GameBoard {
	return &gameBoard{&moveRegister{
		nextMoveBy: move.X,
	}, &BoardRepresentation{}}
}
