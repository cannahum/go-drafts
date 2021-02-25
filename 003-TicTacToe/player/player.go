package player

import "github.com/cannahum/go-drafts/003-TicTacToe/move"

type player struct {
	moveType move.TicTacToeMoveType
}

func (p *player) makeAMove(c move.Coordinates) error {
	return nil
}

func NewPlayer(t move.TicTacToeMoveType) Player {
	return &player{t}
}
