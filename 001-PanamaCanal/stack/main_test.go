package stack

import (
	"testing"

	"github.com/cannahum/go-drafts/001-PanamaCanal/board"
)

func TestPush(t *testing.T) {
	gbs := GameBoardStack{}

	b1 := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}
	gb1 := gbs.Push(&b1)
	if gb1 != &b1 {
		t.Error("Expected to return the same board Got", gb1)
	}
	if gbs.stack[0] != &b1 {
		t.Error("Expected first item to be", b1, "Got", gbs.stack[0])
	}
	if gbs.GetMaxLength() != 1 {
		t.Error("Expected maxLen to be 1 Got", gbs.GetMaxLength())
	}
	if gbs.GetCurrentLength() != 1 {
		t.Error("Expected current length to be 1 Got", gbs.GetCurrentLength())
	}

	b2 := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}
	gb2 := gbs.Push(&b2)
	if gb2 != &b2 {
		t.Error("Expected to return the same board Got", gb2)
	}
	if gbs.stack[0] != &b1 || gbs.stack[1] != &b2 {
		t.Errorf("Expected [%v, %v], Got [%v, %v]\n", &b1, b2, gbs.stack[0], gbs.stack[1])
	}
	if gbs.GetMaxLength() != 2 {
		t.Error("Expected maxLen to be 2 Got", gbs.GetMaxLength())
	}
	if gbs.GetCurrentLength() != 2 {
		t.Error("Expected current length to be 2 Got", gbs.GetCurrentLength())
	}
}

func BenchmarkPush(t *testing.B) {
	gbs := GameBoardStack{}

	b := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	for i := 0; i < t.N; i++ {
		gbs.Push(&b)
	}
}

func TestPop_NonEmpty(t *testing.T) {
	b1 := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	b2 := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	gbs := GameBoardStack{
		stack:  []*board.GameBoard{&b1, &b2},
		maxLen: 2,
	}

	received, err := gbs.Pop()
	if err != nil {
		t.Error("Expected err to be nil Got", err)
	}
	if received != &b2 {
		t.Error("Expected", b2, "Got", received)
	}
	if gbs.GetMaxLength() != 2 {
		t.Error("Expected maxLen to be 2 Got", gbs.GetMaxLength())
	}
	if gbs.GetCurrentLength() != 1 {
		t.Error("Expected current length to be 1 Got", gbs.GetCurrentLength())
	}

	expectedState := GameBoardStack{
		stack: []*board.GameBoard{&b1},
	}

	if len(expectedState.stack) != len(gbs.stack) {
		t.Error("Expected length of gbs to be", len(expectedState.stack), "Got", len(gbs.stack))
	}

	for i, gameBoard := range gbs.stack {
		if gameBoard.Board != gbs.stack[i].Board {
			t.Errorf("expected gbs.stack[%d].Board to be %v\nGot\n%v\n", i, expectedState.stack[i], gameBoard)
		}
	}

	received, err = gbs.Pop()
	if err != nil {
		t.Error("Expected err to be nil Got", err)
	}
	if received != &b1 {
		t.Error("Expected", b1, "Got", received)
	}
	if gbs.GetCurrentLength() != 0 {
		t.Error("Expected current length to be 0 Got", gbs.GetCurrentLength())
	}
}

func TestPop_Empty(t *testing.T) {
	gbs := GameBoardStack{}

	_, err := gbs.Pop()

	if err == nil {
		t.Errorf("expected error Got nil")
	} else if err.Error() != "stack is empty, nothing to pop" {
		t.Errorf("Expected error Got %s", err.Error())
	}

}
