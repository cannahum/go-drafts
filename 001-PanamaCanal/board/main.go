// Package board defines the structure and methods for a single game board
package board

import (
	"errors"
	"log"
)

// Board is a 2D array to hold two lines of 6 characters
type Board [2][6]string

// GameBoard is a struct that holds a single game board.
// It is linked to how it was derived and another board that derives from it.
type GameBoard struct {
	Board             Board
	prev              *GameBoard
	tileMoveDirection string
	tileMoveChar      string
}

// GetChar returns the character (string) of the tile that was moved
func (gb *GameBoard) GetChar() string {
	return gb.tileMoveChar
}

// GetDirection returns the direction of the tile that was moved to achieve this board.
// It'll return "up", "down", "left", or "right"
func (gb *GameBoard) GetDirection() string {
	return gb.tileMoveDirection
}

// GetPrev returns the previous GameBoard from which this board was created
func (gb *GameBoard) GetPrev() *GameBoard {
	return gb.prev
}

// SolutionBoard is the target board, the solution to the Panama Canal puzzle
var SolutionBoard = Board{
	{"P", "A", "N", "A", "M", "A"},
	{"C", "A", "N", "A", "L", ""},
}

func (gb *GameBoard) String() string {
	bb := *gb

	s := ""
	for i, line := range bb.Board {
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
		newBoard.Board[li][ci] = newBoard.Board[li-1][ci]
		newBoard.Board[li-1][ci] = ""
		newBoard.prev = b
		newBoard.tileMoveDirection = "down"
		newBoard.tileMoveChar = newBoard.Board[li][ci]
		result = append(result, &newBoard)
	} else {
		// Create Scenario: moving tile up
		newBoard := currentBoard
		newBoard.Board[li][ci] = newBoard.Board[li+1][ci]
		newBoard.Board[li+1][ci] = ""
		newBoard.prev = b
		newBoard.tileMoveDirection = "up"
		newBoard.tileMoveChar = newBoard.Board[li][ci]
		result = append(result, &newBoard)
	}

	// Create Scenario: moving tile left (if possible)
	if ci != 5 {
		newBoard := currentBoard
		newBoard.Board[li][ci] = newBoard.Board[li][ci+1]
		newBoard.Board[li][ci+1] = ""
		newBoard.prev = b
		newBoard.tileMoveDirection = "left"
		newBoard.tileMoveChar = newBoard.Board[li][ci]
		result = append(result, &newBoard)
	}

	// Create Scenario: moving tile right (if possible)
	if ci != 0 {
		newBoard := currentBoard
		newBoard.Board[li][ci] = newBoard.Board[li][ci-1]
		newBoard.Board[li][ci-1] = ""
		newBoard.prev = b
		newBoard.tileMoveDirection = "right"
		newBoard.tileMoveChar = newBoard.Board[li][ci]
		result = append(result, &newBoard)
	}

	return result
}

func getEmptyTileCoordinates(b *GameBoard) (int, int, error) {
	board := b.Board
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
	return b.Board == SolutionBoard
}
