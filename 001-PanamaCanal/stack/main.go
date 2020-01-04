// Package stack implements a stack (FIFO) data structure holding GameBoard objects
package stack

import (
	"errors"

	"github.com/cannahum/go-drafts/001-PanamaCanal/board"
)

// GameBoardStack is a type to which GameBoards can pushed and popped
type GameBoardStack struct {
	stack  []*board.GameBoard
	maxLen int
}

// GetMaxLength returns the largest size the stack has gotten during the operation
func (gbs *GameBoardStack) GetMaxLength() int {
	return gbs.maxLen
}

// GetCurrentLength returns the the current size of the stack
func (gbs *GameBoardStack) GetCurrentLength() int {
	return len(gbs.stack)
}

// Push adds an item to the stack
func (gbs *GameBoardStack) Push(b *board.GameBoard) *board.GameBoard {
	s := gbs.stack
	gbs.stack = append(s, b)

	length := len(gbs.stack)
	if length > gbs.maxLen {
		gbs.maxLen = length
	}

	return b
}

// Pop removes the last pushed item from the stack
func (gbs *GameBoardStack) Pop() (*board.GameBoard, error) {
	stack := gbs.stack
	lastIndex := len(stack)
	if lastIndex == 0 {
		return nil, errors.New("Stack is empty, nothing to pop")
	}

	poppedValue := stack[lastIndex-1]
	gbs.stack = stack[0 : lastIndex-1]
	return poppedValue, nil
}
