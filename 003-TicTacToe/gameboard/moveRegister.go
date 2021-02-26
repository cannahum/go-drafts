package gameboard

import (
	"errors"

	"github.com/cannahum/go-drafts/003-TicTacToe/move"
)

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
