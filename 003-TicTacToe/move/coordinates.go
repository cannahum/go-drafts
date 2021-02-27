package move

// Coordinates depict the location on a board.
type Coordinates struct {
	Row int
	Col int
}

func NewCoordinates(row, col int) Coordinates {
	return Coordinates{row, col}
}
