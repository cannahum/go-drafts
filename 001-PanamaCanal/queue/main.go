// Package queue allows for a list of potential gameboards to be held in a BFS fashion
package queue

import (
	"errors"

	"github.com/cannahum/go-drafts/001-PanamaCanal/board"
)

// GameBoardQueue is the main struct for the queue
type GameBoardQueue struct {
	queue  []*board.GameBoard
	maxLen int
}

// String prints out the queue
func (gbq GameBoardQueue) String() string {
	l := len(gbq.queue)
	if l == 0 {
		return "[]"
	}

	s := "[\n"
	for i, gameBoard := range gbq.queue {
		b := *gameBoard
		s += b.String()

		if i < l-1 {
			s += "\n\n"
		}
	}
	s += "\n]"

	return s
}

// Enqueue adds a new GameBoard pointer to the end of the queue
func (gbq *GameBoardQueue) Enqueue(b *board.GameBoard) *GameBoardQueue {
	q := gbq.queue
	gbq.queue = append(q, b)
	newLength := len(gbq.queue)
	if newLength > gbq.maxLen {
		gbq.maxLen = newLength
	}
	return gbq
}

// Dequeue adds a new GameBoard pointer to the end of the queue
func (gbq *GameBoardQueue) Dequeue() (*board.GameBoard, error) {
	if len(gbq.queue) == 0 {
		return nil, errors.New("The Queue is empty, nothing to dequeue")
	}
	b := gbq.queue[0]
	gbq.queue = gbq.queue[1:]

	return b, nil
}
