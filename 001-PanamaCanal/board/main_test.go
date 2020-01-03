package board

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetChar(t *testing.T) {
	b := GameBoard{
		Board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
		tileMoveDirection: "left",
		tileMoveChar:      "L",
	}

	char := b.GetChar()
	if char != "L" {
		t.Error("Expected L Got", char)
	}
}
func TestGetDirection(t *testing.T) {
	b := GameBoard{
		Board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
		tileMoveDirection: "left",
		tileMoveChar:      "L",
	}
	direction := b.GetDirection()
	if direction != "left" {
		t.Error("Expected left Got", direction)
	}
}
func TestGetPrev(t *testing.T) {
	b1 := GameBoard{
		Board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
		tileMoveDirection: "left",
		tileMoveChar:      "L",
	}
	b2 := GameBoard{
		Board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
		tileMoveDirection: "left",
		tileMoveChar:      "L",
		prev:              &b1,
	}

	prev := b2.GetPrev()
	if prev != &b1 {
		t.Error("Expected", &b1, "Got", prev)
	}
}

func TestString(t *testing.T) {
	b := GameBoard{
		Board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	expected := fmt.Sprintf("\tP  A  N  A  M  A\n\tC  A  N  A  L  []")
	s := b.String()

	if s != expected {
		t.Errorf("Expected\n%v\nGot\n%v", expected, s)
	}
}

func BenchmarkString(t *testing.B) {
	g := GameBoard{
		Board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	for i := 0; i < t.N; i++ {
		_ = g.String()
	}
}

type getEmptyTileCoordinatesTestCase struct {
	gameBoard GameBoard
	answer    []int
}

func TestGetEmptyTileCoordinates_NoError(t *testing.T) {
	cases := []getEmptyTileCoordinatesTestCase{
		{
			gameBoard: GameBoard{
				Board: Board{
					{"P", "A", "N", "A", "M", "A"},
					{"C", "A", "N", "A", "L", ""},
				},
			},
			answer: []int{1, 5},
		},
		{
			gameBoard: GameBoard{
				Board: Board{
					{"P", "A", "N", "A", "M", ""},
					{"C", "A", "N", "A", "L", "A"},
				},
			},
			answer: []int{0, 5},
		},
		{
			gameBoard: GameBoard{
				Board: Board{
					{"", "P", "A", "N", "A", "M"},
					{"C", "A", "N", "A", "L", "A"},
				},
			},
			answer: []int{0, 0},
		},
		{
			gameBoard: GameBoard{
				Board: Board{
					{"C", "P", "A", "N", "A", "M"},
					{"", "A", "N", "A", "L", "A"},
				},
			},
			answer: []int{1, 0},
		},
		{
			gameBoard: GameBoard{
				Board: Board{
					{"C", "P", "A", "N", "A", "M"},
					{"A", "", "N", "A", "L", "A"},
				},
			},
			answer: []int{1, 1},
		},
	}

	for _, testCase := range cases {
		li, ci, err := getEmptyTileCoordinates(&testCase.gameBoard)
		if err != nil {
			t.Error("Received error:", err)
		}

		if li != testCase.answer[0] {
			t.Error("Expected li:", testCase.answer[0], "Got", li)
		}

		if ci != testCase.answer[1] {
			t.Error("Expected ci:", testCase.answer[1], "Got", ci)
		}
	}
}

func TestGetEmptyTileCoordinates_WithError(t *testing.T) {
	badBoard := GameBoard{
		Board: Board{
			{"X", "P", "A", "N", "A", "M"},
			{"C", "A", "N", "A", "L", "A"},
		},
	}

	li, ci, err := getEmptyTileCoordinates(&badBoard)

	if err == nil {
		t.Error("Expected error Got nil")
	}

	if li != -1 || ci != -1 {
		t.Error("Expected both coordinates to be -1. Got", li, ci)
	}
}

func BenchmarkGetEmptyTileCoordinates(t *testing.B) {
	g := GameBoard{
		Board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	for i := 0; i < t.N; i++ {
		_, _, _ = getEmptyTileCoordinates(&g)
	}
}

type variateTestCase struct {
	gameBoard GameBoard
	answer    []GameBoard
}

func TestVariate(t *testing.T) {
	canamaPanalStart := GameBoard{
		Board: Board{
			{"C", "A", "N", "A", "M", "A"},
			{"P", "A", "N", "A", "L", ""},
		},
	}

	canamaPanalDown := GameBoard{
		Board: Board{
			{"C", "A", "N", "A", "M", ""},
			{"P", "A", "N", "A", "L", "A"},
		},
	}

	canamaPanalRight := GameBoard{
		Board: Board{
			{"C", "A", "N", "A", "M", "A"},
			{"P", "A", "N", "A", "", "L"},
		},
	}

	canamaPanalDownRight := GameBoard{
		Board: Board{
			{"C", "A", "N", "A", "", "M"},
			{"P", "A", "N", "A", "L", "A"},
		},
	}

	canamaPanalNorthWest := GameBoard{
		Board: Board{
			{"", "C", "A", "N", "A", "M"},
			{"P", "A", "N", "A", "L", "A"},
		},
	}

	canamaPanalNorthWestLeft := GameBoard{
		Board: Board{
			{"C", "", "A", "N", "A", "M"},
			{"P", "A", "N", "A", "L", "A"},
		},
	}

	canamaPanalNorthWestUp := GameBoard{
		Board: Board{
			{"P", "C", "A", "N", "A", "M"},
			{"", "A", "N", "A", "L", "A"},
		},
	}

	testCases := []variateTestCase{
		{
			gameBoard: canamaPanalStart,
			answer: []GameBoard{
				canamaPanalDown,
				canamaPanalRight,
			},
		},
		{
			gameBoard: canamaPanalDown,
			answer: []GameBoard{
				canamaPanalDownRight,
				canamaPanalStart,
			},
		},
		{
			gameBoard: canamaPanalNorthWest,
			answer: []GameBoard{
				canamaPanalNorthWestLeft,
				canamaPanalNorthWestUp,
			},
		},
	}

	for _, scenario := range testCases {
		result := Variate(&scenario.gameBoard)

		if len(result) != len(scenario.answer) {
			t.Errorf("Expected %d results, Got %d\n", len(scenario.answer), len(result))
		}

		for _, variationBoard := range result {
			foundIdentical := false
			for _, expectedBoard := range scenario.answer {
				if variationBoard.Board == expectedBoard.Board {
					foundIdentical = true
				}
			}
			if !foundIdentical {
				t.Error("Expected", scenario.answer)
			}

			// Check that prev field is set
			typeOfPrev := reflect.TypeOf(variationBoard.prev).String()
			if typeOfPrev != "*board.GameBoard" {
				t.Error("Expected reference to prev but NOT ok")
			}
		}
	}
}

func TestIsSolutionBoard(t *testing.T) {
	truthy := GameBoard{
		Board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	res := IsSolutionBoard(&truthy)
	if !res {
		t.Error("Expected true Got false")
	}

	falsy := GameBoard{
		Board: Board{
			{"C", "A", "N", "A", "M", "A"},
			{"P", "A", "N", "A", "L", ""},
		},
	}
	res = IsSolutionBoard(&falsy)
	if res {
		t.Error("Expected false Got true")
	}
}

func BenchmarkIsSolutionBoard_Truthy(t *testing.B) {
	truthy := GameBoard{
		Board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	for i := 0; i < t.N; i++ {
		_ = IsSolutionBoard(&truthy)
	}
}

func BenchmarkIsSolutionBoard_Falsy(t *testing.B) {
	falsy := GameBoard{
		Board: Board{
			{"C", "A", "N", "A", "M", "A"},
			{"P", "A", "N", "A", "L", ""},
		},
	}

	for i := 0; i < t.N; i++ {
		_ = IsSolutionBoard(&falsy)
	}
}
