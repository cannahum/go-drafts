package gameboard

// GameBoard interface is a way to interact with an existing board instance
type GameBoard interface {
	RegisterMove(MoveInGame) (bool, error)
	GetCurrentBoardState() BoardRepresentation
	GetBoardStateAt(int) BoardRepresentation
}
