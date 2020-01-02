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
		g.String()
	}
}
