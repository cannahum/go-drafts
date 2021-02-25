package player

type player struct {
	moveType move.TicTacToeMoveType
}

func NewPlayer(t move.TicTacToeMoveType) Player {
	return player{t}
}
