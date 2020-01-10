package queue

import (
	"fmt"
	"testing"

	"github.com/cannahum/go-drafts/001-PanamaCanal/board"
)

func TestGetCurrentLength(t *testing.T) {
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

	gbq := GameBoardQueue{
		queue: []*board.GameBoard{},
	}

	x := gbq.GetCurrentLength()
	if x != 0 {
		t.Error("Expected 0 Got", x)
	}

	gbq.Enqueue(&b1)
	gbq.Enqueue(&b2)

	x = gbq.GetCurrentLength()
	if x != 2 {
		t.Error("Expected 2 Got", x)
	}

	_, _ = gbq.Dequeue()
	x = gbq.GetCurrentLength()
	if x != 1 {
		t.Error("Expected 1 Got", x)
	}
}

func TestGetMaxLength(t *testing.T) {
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

	gbq := GameBoardQueue{
		queue: []*board.GameBoard{},
	}

	x := gbq.GetMaxLength()
	if x != 0 {
		t.Error("Expected 0 Got", x)
	}

	gbq.Enqueue(&b1)
	gbq.Enqueue(&b2)

	x = gbq.GetMaxLength()
	if x != 2 {
		t.Error("Expected 2 Got", x)
	}

	_, _ = gbq.Dequeue()
	x = gbq.GetMaxLength()
	if x != 2 {
		t.Error("Expected 2 Got", x)
	}
}

func TestString_NotEmpty(t *testing.T) {
	b := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	gbq := GameBoardQueue{
		queue: []*board.GameBoard{
			&b,
		},
	}

	expected := fmt.Sprintf("[\n\tP  A  N  A  M  A\n\tC  A  N  A  L  []\n]")
	s := gbq.String()

	if s != expected {
		t.Errorf("Expected\n%v\nGot\n%v", expected, s)
	}
}

func TestString_Empty(t *testing.T) {
	gbq := GameBoardQueue{
		queue: []*board.GameBoard{},
	}

	expected := fmt.Sprintf("[]")
	s := gbq.String()

	if s != expected {
		t.Errorf("Expected\n%v\nGot\n%v", expected, s)
	}
}

func BenchmarkString(t *testing.B) {
	b := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	gbq := GameBoardQueue{
		queue: []*board.GameBoard{
			&b,
		},
	}

	for i := 0; i < t.N; i++ {
		_ = gbq.String()
	}
}

func TestEnqueue_Empty(t *testing.T) {
	b1 := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}
	gbq := GameBoardQueue{
		queue: []*board.GameBoard{},
	}
	gbq.Enqueue(&b1)

	if gbq.queue[0] != &b1 {
		t.Errorf("Expected [%v], Got [%v]\n", &b1, gbq.queue[0])
	}

	if gbq.maxLen != 1 {
		t.Error("Expected length to go up by one, Got", gbq.maxLen)
	}
}

func TestEnqueue_NonEmpty(t *testing.T) {
	b1 := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	gbq := GameBoardQueue{
		queue: []*board.GameBoard{
			&b1,
		},
		maxLen: 1,
	}

	oldLength := gbq.maxLen

	b2 := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	gbq.Enqueue(&b2)

	if gbq.queue[0] != &b1 || gbq.queue[1] != &b2 {
		t.Errorf("Expected [%v, %v], Got [%v, %v]\n", &b1, b2, gbq.queue[0], gbq.queue[1])
	}

	if gbq.maxLen != oldLength+1 {
		t.Error("Expected length to go up by one, Got", gbq.maxLen)
	}
}

func BenchmarkEnqueue(t *testing.B) {
	gbq := GameBoardQueue{
		queue: []*board.GameBoard{},
	}

	b := board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	}

	for i := 0; i < t.N; i++ {
		gbq.Enqueue(&b)
	}
}

func TestDequeue_NonEmpty(t *testing.T) {
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

	gbq := GameBoardQueue{
		queue: []*board.GameBoard{
			&b1,
			&b2,
		},
	}

	oldLength := gbq.maxLen

	received, err := gbq.Dequeue()
	if err != nil {
		t.Error("Expected err to be nil Got", err)
	}

	if received != &b1 {
		t.Error("Expected b1's address Got", received)
	}

	expectedState := GameBoardQueue{
		queue: []*board.GameBoard{
			&b2,
		},
		maxLen: 2,
	}

	if len(expectedState.queue) != len(gbq.queue) {
		t.Error("Expected length of gbq to be", len(expectedState.queue), "Got", len(gbq.queue))
	}

	for i, gameBoard := range gbq.queue {
		if gameBoard.Board != gbq.queue[i].Board {
			t.Errorf("Expected gbq.queue[%d].Board to be %v\nGot\n%v\n", i, expectedState.queue[i], gameBoard)
		}
	}

	if gbq.maxLen != oldLength {
		t.Error("Expected maxLen to not change, Got", gbq.maxLen)
	}
}

func TestDequeue_Empty(t *testing.T) {
	gbq := GameBoardQueue{
		queue: []*board.GameBoard{},
	}

	_, err := gbq.Dequeue()

	if err == nil {
		t.Errorf("Expected error Got nil")
	} else if err.Error() != "queue is empty, nothing to dequeue" {
		t.Errorf("Expected error Got %s", err.Error())
	}
}
