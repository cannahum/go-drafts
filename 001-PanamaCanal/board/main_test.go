package board

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	b := GameBoard{
		board: Board{
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
		board: Board{
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
				board: Board{
					{"P", "A", "N", "A", "M", "A"},
					{"C", "A", "N", "A", "L", ""},
				},
			},
			answer: []int{1, 5},
		},
		{
			gameBoard: GameBoard{
				board: Board{
					{"P", "A", "N", "A", "M", ""},
					{"C", "A", "N", "A", "L", "A"},
				},
			},
			answer: []int{0, 5},
		},
		{
			gameBoard: GameBoard{
				board: Board{
					{"", "P", "A", "N", "A", "M"},
					{"C", "A", "N", "A", "L", "A"},
				},
			},
			answer: []int{0, 0},
		},
		{
			gameBoard: GameBoard{
				board: Board{
					{"C", "P", "A", "N", "A", "M"},
					{"", "A", "N", "A", "L", "A"},
				},
			},
			answer: []int{1, 0},
		},
		{
			gameBoard: GameBoard{
				board: Board{
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
		board: Board{
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
		board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	for i := 0; i < t.N; i++ {
		_, _, _ = getEmptyTileCoordinates(&g)
	}
}

func TestIsSolutionBoard(t *testing.T) {
	truthy := GameBoard{
		board: Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	res := IsSolutionBoard(&truthy)
	if !res {
		t.Error("Expected true Got false")
	}

	falsy := GameBoard{
		board: Board{
			{"C", "A", "N", "A", "M", "A"},
			{"P", "A", "N", "A", "L", ""},
		},
	}
	res = IsSolutionBoard(&falsy)
	if res {
		t.Error("Expected false Got true")
	}
}
