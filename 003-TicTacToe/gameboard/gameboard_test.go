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
	board.RegisterMove(move.X, move.Coordinates{
		Row: 0,
		Col: 0,
	})
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
	board.RegisterMove(move.X, move.Coordinates{
		Row: 0,
		Col: 0,
	})
	board.RegisterMove(move.O, move.Coordinates{
		Row: 1,
		Col: 1,
	})
	board.RegisterMove(move.X, move.Coordinates{
		Row: 2,
		Col: 2,
	})
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
	_, err := board.RegisterMove(move.O, move.Coordinates{
		Row: 0,
		Col: 0,
	})
	if err == nil {
		t.Error("Expected error but didn't get one")
	}

	_, err1 := board.RegisterMove(move.X, move.Coordinates{
		Row: 0,
		Col: 0,
	})

	if err1 != nil {
		t.Error("Expected ok but got error")
	}

	_, err2 := board.RegisterMove(move.X, move.Coordinates{
		Row: 1,
		Col: 1,
	})

	if err2 == nil {
		t.Error("Expected error but didn't get one")
	}
}

func TestNewGameBoardInvalidMoveByWrongCoordinates(t *testing.T) {
	board := NewGameBoard()
	_, err1 := board.RegisterMove(move.X, move.Coordinates{
		Row: 0,
		Col: 0,
	})

	if err1 != nil {
		t.Error("Expected ok but got error")
	}

	_, err2 := board.RegisterMove(move.O, move.Coordinates{
		Row: 0,
		Col: 0,
	})

	if err2 == nil {
		t.Error("Expected error but didn't get one")
	}
}
