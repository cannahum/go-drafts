package gameboard

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cannahum/go-drafts/003-TicTacToe/move"
)

type MoveInGame struct {
	MoveType    move.TicTacToeMoveType
	Coordinates move.Coordinates
}

type BoardRepresentation [3][3]move.TicTacToeMoveType

func (b BoardRepresentation) String() string {
	writeMove := func(m move.TicTacToeMoveType) string {
		switch m {
		case move.O:
			return "O"
		case move.X:
			return "X"
		default:
			return " "
		}
	}
	var sb strings.Builder
	for i, row := range b {
		if i == 1 {
			sb.WriteString(strings.Repeat("-", 11))
			sb.WriteString("\n")
		}

		for j, col := range row {
			if j == 1 {
				sb.WriteString("|")
			}
			sb.WriteString(fmt.Sprintf(" %s ", writeMove(col)))
			if j == 1 {
				sb.WriteString("|")
			}
		}

		if i == 1 {
			sb.WriteString("\n")
			sb.WriteString(strings.Repeat("-", 11))
		}

		if i < 2 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (b BoardRepresentation) hasTicTacToeAtCoordinates(coorA, coorB, coorC move.Coordinates) bool {
	moveA := b[coorA.Row][coorA.Col]
	moveB := b[coorB.Row][coorB.Col]
	moveC := b[coorC.Row][coorC.Col]
	return moveA == moveB && moveB == moveC
}

func (b BoardRepresentation) hasRowTicTacToe() bool {
	return b.hasTicTacToeAtCoordinates(move.Coordinates{0, 0}, move.Coordinates{0, 1}, move.Coordinates{0, 2}) ||
		b.hasTicTacToeAtCoordinates(move.Coordinates{1, 0}, move.Coordinates{1, 1}, move.Coordinates{1, 2}) ||
		b.hasTicTacToeAtCoordinates(move.Coordinates{2, 0}, move.Coordinates{2, 1}, move.Coordinates{2, 2})
}

func (b BoardRepresentation) hasColTicTacToe() bool {
	return b.hasTicTacToeAtCoordinates(move.Coordinates{0, 0}, move.Coordinates{1, 0}, move.Coordinates{2, 0}) ||
		b.hasTicTacToeAtCoordinates(move.Coordinates{0, 1}, move.Coordinates{1, 1}, move.Coordinates{2, 1}) ||
		b.hasTicTacToeAtCoordinates(move.Coordinates{0, 2}, move.Coordinates{1, 2}, move.Coordinates{2, 2})
}

func (b BoardRepresentation) hasDiagonalTicTacToe() bool {
	return b.hasTicTacToeAtCoordinates(move.Coordinates{0, 0}, move.Coordinates{1, 1}, move.Coordinates{2, 2}) ||
		b.hasTicTacToeAtCoordinates(move.Coordinates{0, 2}, move.Coordinates{1, 1}, move.Coordinates{2, 0})
}

type moveRegister struct {
	moves      []MoveInGame
	nextMoveBy move.TicTacToeMoveType
}

func (m *moveRegister) isValidMove(newMove MoveInGame) bool {
	if newMove.MoveType != m.nextMoveBy {
		return false
	}

	for _, eachMove := range m.moves {
		if eachMove.Coordinates == newMove.Coordinates {
			return false
		}
	}
	return true
}

func (m *moveRegister) registerMove(newMove MoveInGame) error {
	currentListOfMoves := m.moves

	if !m.isValidMove(newMove) {
		return errors.New("Invalid move")
	}

	currentListOfMoves = append(currentListOfMoves, newMove)
	m.moves = currentListOfMoves
	if newMove.MoveType == move.O {
		m.nextMoveBy = move.X
	} else {
		m.nextMoveBy = move.O
	}
	return nil
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
