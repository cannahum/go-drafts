package gameboard

import (
	"testing"

	"github.com/cannahum/go-drafts/003-TicTacToe/move"
)

func TestNewGameBoard(t *testing.T) {
	board := NewGameBoard()
	emptyBoard := BoardRepresentation{}
	if board.GetCurrentBoardState() != emptyBoard {
		t.Error("Expected new board GetCurrentBoardState to equal zero state board representation but got", board.GetCurrentBoardState())
	}

	if board.GetBoardStateAt(0) != emptyBoard {
		t.Error("Expected new board GetBoardStateAt(0) to equal zero state board representation but got", board.GetBoardStateAt(0))
	}
}

func TestNewGameBoardPrint(t *testing.T) {
	emptyBoard := BoardRepresentation{}
	boardPrint := emptyBoard.String()
	expectedPrint := `   |   |   
-----------
   |   |   
-----------
   |   |   `

	if boardPrint != expectedPrint {
		t.Errorf("Expected \n%s\n but got \n%s\n", expectedPrint, boardPrint)
	}
}

func TestNewGameBoardMakeAMove(t *testing.T) {
	board := NewGameBoard()
	board.RegisterMove(xMakesMove(0, 0))
	currentBoardState := board.GetCurrentBoardState()
	currentPrint := currentBoardState.String()

	expectedBoardPrint := ` X |   |   
-----------
   |   |   
-----------
   |   |   `

	if currentPrint != expectedBoardPrint {
		t.Errorf("Expected \n%s\n but got \n%s\n", expectedBoardPrint, currentPrint)
	}
}

func TestNewGameBoardMakeMoves(t *testing.T) {
	board := NewGameBoard()
	board.RegisterMove(xMakesMove(0, 0))
	board.RegisterMove(oMakesMove(1, 1))
	board.RegisterMove(xMakesMove(2, 2))
	currentBoardState := board.GetCurrentBoardState()
	currentPrint := currentBoardState.String()

	expectedBoardPrint := ` X |   |   
-----------
   | O |   
-----------
   |   | X `

	if currentPrint != expectedBoardPrint {
		t.Errorf("Expected \n%s\n but got \n%s\n", expectedBoardPrint, currentPrint)
	}

	// Test at a certain point in time
	boardAtState2 := board.GetBoardStateAt(2)
	state2Print := boardAtState2.String()
	expectedState2Print := ` X |   |   
-----------
   | O |   
-----------
   |   |   `

	if state2Print != expectedState2Print {
		t.Errorf("At State 2, expected \n%s\n but got \n%s\n", expectedBoardPrint, currentPrint)
	}
}

func TestNewGameBoardInvalidMoveByWrongPlayer(t *testing.T) {
	board := NewGameBoard()
	_, err := board.RegisterMove(oMakesMove(0, 0))
	if err == nil {
		t.Error("Expected error but didn't get one")
	}

	_, err1 := board.RegisterMove(xMakesMove(0, 0))
	if err1 != nil {
		t.Error("Expected ok but got error")
	}

	_, err2 := board.RegisterMove(xMakesMove(1, 1))
	if err2 == nil {
		t.Error("Expected error but didn't get one")
	}
}

func TestNewGameBoardInvalidMoveByWrongCoordinates(t *testing.T) {
	board := NewGameBoard()
	_, err1 := board.RegisterMove(xMakesMove(0, 0))
	if err1 != nil {
		t.Error("Expected ok but got error")
	}

	_, err2 := board.RegisterMove(oMakesMove(0, 0))
	if err2 == nil {
		t.Error("Expected error but didn't get one")
	}
}

func TestGameEndingWithRows(t *testing.T) {
	board1 := NewGameBoard()
	board1.RegisterMove(xMakesMove(0, 0))
	board1.RegisterMove(oMakesMove(1, 1))
	board1.RegisterMove(xMakesMove(0, 1))
	board1.RegisterMove(oMakesMove(2, 0))
	gameFinished, _ := board1.RegisterMove(xMakesMove(0, 2))

	if gameFinished != true {
		t.Error("Expected game end but got false")
	}

	board2 := NewGameBoard()
	board2.RegisterMove(xMakesMove(0, 0))
	board2.RegisterMove(oMakesMove(1, 1))
	board2.RegisterMove(xMakesMove(2, 2))
	board2.RegisterMove(oMakesMove(1, 0))
	board2.RegisterMove(xMakesMove(0, 2))
	gameFinished, _ = board2.RegisterMove(oMakesMove(1, 2))

	if gameFinished != true {
		t.Error("Expected game end but got false")
	}

	board3 := NewGameBoard()
	board3.RegisterMove(xMakesMove(2, 0))
	board3.RegisterMove(oMakesMove(1, 1))
	board3.RegisterMove(xMakesMove(2, 1))
	board3.RegisterMove(oMakesMove(0, 2))
	gameFinished, _ = board3.RegisterMove(xMakesMove(2, 2))
	if gameFinished != true {
		t.Error("Expected game end but got false")
	}
}

func TestGameEndingWithCols(t *testing.T) {
	board1 := NewGameBoard()
	board1.RegisterMove(xMakesMove(0, 0))
	board1.RegisterMove(oMakesMove(1, 1))
	board1.RegisterMove(xMakesMove(1, 0))
	board1.RegisterMove(oMakesMove(1, 2))
	gameFinished, _ := board1.RegisterMove(xMakesMove(2, 0))

	if gameFinished != true {
		t.Error("Expected game end but got false")
	}

	board2 := NewGameBoard()
	board2.RegisterMove(xMakesMove(0, 0))
	board2.RegisterMove(oMakesMove(1, 1))
	board2.RegisterMove(xMakesMove(2, 2))
	board2.RegisterMove(oMakesMove(0, 1))
	board2.RegisterMove(xMakesMove(0, 2))
	gameFinished, _ = board2.RegisterMove(oMakesMove(2, 1))

	if gameFinished != true {
		t.Error("Expected game end but got false")
	}

	board3 := NewGameBoard()
	board3.RegisterMove(xMakesMove(0, 2))
	board3.RegisterMove(oMakesMove(1, 1))
	board3.RegisterMove(xMakesMove(1, 2))
	board3.RegisterMove(oMakesMove(0, 1))
	gameFinished, _ = board3.RegisterMove(xMakesMove(2, 2))
	if gameFinished != true {
		t.Error("Expected game end but got false")
	}
}

func TestGameEndingWithDiagonals(t *testing.T) {
	board1 := NewGameBoard()
	board1.RegisterMove(xMakesMove(0, 0))
	board1.RegisterMove(oMakesMove(0, 2))
	board1.RegisterMove(xMakesMove(1, 1))
	board1.RegisterMove(oMakesMove(1, 2))
	gameFinished, _ := board1.RegisterMove(xMakesMove(2, 2))

	if gameFinished != true {
		t.Error("Expected game end but got false")
	}

	board2 := NewGameBoard()
	board2.RegisterMove(xMakesMove(0, 0))
	board2.RegisterMove(oMakesMove(0, 2))
	board2.RegisterMove(xMakesMove(2, 2))
	board2.RegisterMove(oMakesMove(1, 1))
	board2.RegisterMove(xMakesMove(0, 1))
	gameFinished, _ = board2.RegisterMove(oMakesMove(2, 0))

	if gameFinished != true {
		t.Error("Expected game end but got false")
	}
}

func getCoordinates(row, col int) move.Coordinates {
	return move.Coordinates{
		Row: row,
		Col: col,
	}
}

func xMakesMove(row, col int) MoveInGame {
	return MoveInGame{
		MoveType:    move.X,
		Coordinates: getCoordinates(row, col),
	}
}

func oMakesMove(row, col int) MoveInGame {
	return MoveInGame{
		MoveType:    move.O,
		Coordinates: getCoordinates(row, col),
	}
}
