package gameboard

import (
	"fmt"
	"strings"

	"github.com/cannahum/go-drafts/003-TicTacToe/move"
)

var (
	firstRow          = [3]move.Coordinates{move.NewCoordinates(0, 0), move.NewCoordinates(0, 1), move.NewCoordinates(0, 2)}
	secondRow         = [3]move.Coordinates{move.NewCoordinates(1, 0), move.NewCoordinates(1, 1), move.NewCoordinates(1, 2)}
	thirdRow          = [3]move.Coordinates{move.NewCoordinates(2, 0), move.NewCoordinates(2, 1), move.NewCoordinates(2, 2)}
	firstColumn       = [3]move.Coordinates{move.NewCoordinates(0, 0), move.NewCoordinates(1, 0), move.NewCoordinates(2, 0)}
	secondColumn      = [3]move.Coordinates{move.NewCoordinates(0, 1), move.NewCoordinates(1, 1), move.NewCoordinates(2, 1)}
	thirdColumn       = [3]move.Coordinates{move.NewCoordinates(0, 2), move.NewCoordinates(1, 2), move.NewCoordinates(2, 2)}
	northWestDiagonal = [3]move.Coordinates{move.NewCoordinates(0, 0), move.NewCoordinates(1, 1), move.NewCoordinates(2, 2)}
	northEastDiagonal = [3]move.Coordinates{move.NewCoordinates(0, 2), move.NewCoordinates(1, 1), move.NewCoordinates(2, 0)}
)

type BoardRepresentation [3][3]move.TicTacToeMoveType
type WinnerCoordinates [3]move.Coordinates

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
	return moveA != 0 && moveA == moveB && moveB == moveC
}

func (b BoardRepresentation) findRowTicTacToe() *WinnerCoordinates {
	var winner WinnerCoordinates
	hasWinner := false
	if b.hasTicTacToeAtCoordinates(firstRow[0], firstRow[1], firstRow[2]) {
		winner = firstRow
		hasWinner = true
	}

	if b.hasTicTacToeAtCoordinates(secondRow[0], secondRow[1], secondRow[2]) {
		winner = secondRow
		hasWinner = true
	}

	if b.hasTicTacToeAtCoordinates(thirdRow[0], thirdRow[1], thirdRow[2]) {
		winner = thirdRow
		hasWinner = true
	}

	if hasWinner {
		return &winner
	}
	return nil
}

func (b BoardRepresentation) findColTicTacToe() *WinnerCoordinates {
	var winner WinnerCoordinates
	hasWinner := false
	if b.hasTicTacToeAtCoordinates(firstColumn[0], firstColumn[1], firstColumn[2]) {
		winner = firstColumn
		hasWinner = true
	}

	if b.hasTicTacToeAtCoordinates(secondColumn[0], secondColumn[1], secondColumn[2]) {
		winner = secondColumn
		hasWinner = true
	}

	if b.hasTicTacToeAtCoordinates(thirdColumn[0], thirdColumn[1], thirdColumn[2]) {
		winner = thirdColumn
		hasWinner = true
	}

	if hasWinner {
		return &winner
	}
	return nil
}

func (b BoardRepresentation) findDiagonalTicTacToe() *WinnerCoordinates {
	var winner WinnerCoordinates
	hasWinner := false
	if b.hasTicTacToeAtCoordinates(northWestDiagonal[0], northWestDiagonal[1], northWestDiagonal[2]) {
		winner = northWestDiagonal
		hasWinner = true
	}

	if b.hasTicTacToeAtCoordinates(northEastDiagonal[0], northEastDiagonal[1], northEastDiagonal[2]) {
		winner = northEastDiagonal
		hasWinner = true
	}

	if hasWinner {
		return &winner
	}
	return nil
}

func (b BoardRepresentation) hasRowTicTacToe() bool {
	rowWinner := b.findRowTicTacToe()
	// memoize, or take note of this
	return rowWinner != nil
}

func (b BoardRepresentation) hasColTicTacToe() bool {
	rowWinner := b.findColTicTacToe()
	// memoize, or take note of this
	return rowWinner != nil
}

func (b BoardRepresentation) hasDiagonalTicTacToe() bool {
	rowWinner := b.findDiagonalTicTacToe()
	// memoize, or take note of this
	return rowWinner != nil
}
