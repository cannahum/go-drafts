// Package board defines the structure and methods for a single game board
package board

import "errors"

import "log"

// GameBoard is a struct that holds a single game board.
// It is linked to how it was derived and another board that derives from it.
type GameBoard struct {
	board Board
	prev  *GameBoard
	next  *GameBoard
}

// Board is a 2D array to hold two lines of 6 characters
type Board [2][6]string

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

// Variate is a method that generates potential game boards by moving tiles legally
// and returns the potentials as an array of GameBoard pointers
func (b *GameBoard) Variate() []*GameBoard {
	result := []*GameBoard{}

	li, ci, err := getEmptyTileCoordinates(b)
	if err != nil {
		log.Fatal("Bad board generated in", *b)
	}

	if li == 1 {
		// Create Scenario: moving tile down (if possible)

	} else {
		// Create Scenario: moving tile up

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
