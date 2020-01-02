// Package board defines the structure and methods for a single game board
package board

import (
	"errors"
	"log"
)

// GameBoard is a struct that holds a single game board.
// It is linked to how it was derived and another board that derives from it.
type GameBoard struct {
	board             Board
	prev              *GameBoard
	tileMoveDirection string
}

// Board is a 2D array to hold two lines of 6 characters
type Board [2][6]string

// SolutionBoard is the target board, the solution to the Panama Canal puzzle
var SolutionBoard = Board{
	{"P", "A", "N", "A", "M", "A"},
	{"C", "A", "N", "A", "L", ""},
}

func (b *GameBoard) String() string {
	bb := *b

	s := ""
	for i, line := range bb.board {
		for j, char := range line {
			if j == 0 {
				s += "\t"
			}

			if char == "" {
				s += "[]"
			} else {
				s += char
			}

			if i == 0 && j == 5 {
				s += "\n"
			} else if j < 5 {
				s += "  "
			}
		}
	}

	return s
}

// Variate that generates potential game boards by moving tiles legally
// and returns the potentials as an array of GameBoard pointers
func Variate(b *GameBoard) []*GameBoard {
	result := []*GameBoard{}

	li, ci, err := getEmptyTileCoordinates(b)
	if err != nil {
		log.Fatal("Bad board generated in", *b)
	}

	currentBoard := *b
	if li == 1 {
		// Create Scenario: moving tile down (if possible)
		newBoard := currentBoard
		newBoard.board[li][ci] = newBoard.board[li-1][ci]
		newBoard.board[li-1][ci] = ""
		newBoard.prev = b
		newBoard.tileMoveDirection = "down"
		result = append(result, &newBoard)
	} else {
		// Create Scenario: moving tile up
		newBoard := currentBoard
		newBoard.board[li][ci] = newBoard.board[li+1][ci]
		newBoard.board[li+1][ci] = ""
		newBoard.prev = b
		newBoard.tileMoveDirection = "up"
		result = append(result, &newBoard)
	}

	// Create Scenario: moving tile left (if possible)
	if ci != 5 {
		newBoard := currentBoard
		newBoard.board[li][ci] = newBoard.board[li][ci+1]
		newBoard.board[li][ci+1] = ""
		newBoard.prev = b
		newBoard.tileMoveDirection = "left"
		result = append(result, &newBoard)
	}

	// Create Scenario: moving tile right (if possible)
	if ci != 0 {
		newBoard := currentBoard
		newBoard.board[li][ci] = newBoard.board[li][ci-1]
		newBoard.board[li][ci-1] = ""
		newBoard.prev = b
		newBoard.tileMoveDirection = "right"
		result = append(result, &newBoard)
	}

	return result
}

func getEmptyTileCoordinates(b *GameBoard) (int, int, error) {
	board := b.board
	for i, line := range board {
		for j, char := range line {
			if char == "" {
				return i, j, nil
			}
		}
	}

	return -1, -1, errors.New("Could not find the empty tile coordinate for")
}

// IsSolutionBoard checks if the board is the target Board
func IsSolutionBoard(b *GameBoard) bool {
	return b.board == SolutionBoard
}
