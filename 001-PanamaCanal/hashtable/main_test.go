package hashtable

import (
	"testing"

	"github.com/cannahum/go-drafts/001-PanamaCanal/board"
)

var boards = []board.GameBoard{
	board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", "A"},
			{"C", "A", "N", "A", "L", ""},
		},
	},
	board.GameBoard{
		Board: board.Board{
			{"P", "A", "N", "A", "M", ""},
			{"C", "A", "N", "A", "L", "A"},
		},
	},
	board.GameBoard{
		Board: board.Board{
			{"", "P", "A", "N", "A", "M"},
			{"C", "A", "N", "A", "L", "A"},
		},
	},
	board.GameBoard{
		Board: board.Board{
			{"C", "P", "A", "N", "A", "M"},
			{"", "A", "N", "A", "L", "A"},
		},
	},
}

type mockHashComputer struct {
	index        int
	stayConstant bool
}

func (h *mockHashComputer) GetHashKey(b *board.GameBoard) int {
	i := h.index
	if !h.stayConstant {
		h.index = h.index + 1
	}
	return i
}

func TestGetNumOfKeys(t *testing.T) {
	gbt := NewGameBoardHashTable()
	// create mock hash function that returns a different key each time
	hashComputer := mockHashComputer{}

	for _, b := range boards {
		gbt.Insert(&b, &hashComputer)
	}

	x := gbt.GetNumOfKeys()
	if x != len(boards) {
		t.Error("Expected", len(boards), "Got", x)
	}
}

func TestGetNumOfLongestLinkedList_DifferentKeys(t *testing.T) {
	gbt := NewGameBoardHashTable()
	// create mock hash function that returns a different key each time
	hashComputer := mockHashComputer{}

	for _, b := range boards {
		gbt.Insert(&b, &hashComputer)
	}

	x := gbt.GetNumOfLongestLinkedList()
	if x != 1 {
		t.Error("Expected 1 Got", x)
	}
}

func TestGetNumOfLongestLinkedList_ConstantKeys(t *testing.T) {
	gbt := NewGameBoardHashTable()
	// create mock hash function that returns the same key every time
	hashComputer := mockHashComputer{
		stayConstant: true,
	}

	gbt.Insert(&boards[0], &hashComputer)
	gbt.Insert(&boards[1], &hashComputer)
	gbt.Insert(&boards[2], &hashComputer)
	gbt.Insert(&boards[3], &hashComputer)

	x := gbt.GetNumOfLongestLinkedList()
	if x != len(boards) {
		t.Error("Expected", len(boards), "Got", x)
	}
}

func TestHas_Empty(t *testing.T) {
	gbt := NewGameBoardHashTable()

	// create mock hash function
	hashComputer := mockHashComputer{}

	for _, v := range boards {
		if gbt.Has(&v, &hashComputer) {
			t.Error("Expected", v, "not to exist but it does")
		}
	}
}

func TestHas_NotEmptySingleNode(t *testing.T) {
	gbt := NewGameBoardHashTable()

	gbt.table[0] = &hashLinkedListNode{
		gameBoard: &boards[0],
	}

	gbt.table[1] = &hashLinkedListNode{
		gameBoard: &boards[1],
	}

	gbt.table[2] = &hashLinkedListNode{
		gameBoard: &boards[2],
	}

	gbt.table[3] = &hashLinkedListNode{
		gameBoard: &boards[3],
	}

	// create mock hash function
	hashComputer := mockHashComputer{}

	for _, v := range boards {
		result := gbt.Has(&v, &hashComputer)
		if !result {
			t.Error("Expected", v, "to exist but it does not", result)
		}
	}
}

func TestHas_NotEmptyLinkedList(t *testing.T) {
	gbt := NewGameBoardHashTable()

	gbt.table[0] = &hashLinkedListNode{
		gameBoard: &boards[0],
	}

	gbt.table[0].next = &hashLinkedListNode{
		gameBoard: &boards[1],
	}

	gbt.table[0].next.next = &hashLinkedListNode{
		gameBoard: &boards[2],
	}

	// create mock hash function
	hashComputer := mockHashComputer{
		stayConstant: true,
	}

	// Everything should be found, except the last
	for i, v := range boards {
		result := gbt.Has(&v, &hashComputer)

		if i < 3 {
			if !result {
				t.Error("Expected", v, "to exist but it does not")
			}
		} else {
			if result {
				t.Error("Expected", v, "not to exist but it does")
			}
		}
	}
}

func TestInsert_Empty(t *testing.T) {
	gbt := NewGameBoardHashTable()

	// create mock hash function
	hashComputer := mockHashComputer{
		stayConstant: true,
	}
	board0 := boards[0]
	gb, ok := gbt.Insert(&board0, &hashComputer)

	if !ok {
		t.Error("Expected to insert gameboard but NOT ok")
	}

	if gb != &board0 {
		t.Error("Expected to receive", &board0, "Got", gb)
	}

	if gbt.numOfKeys != 1 {
		t.Error("Unexpected number of keys", gbt.numOfKeys)
	}

	if gbt.numOfLongestLinkedList != 1 {
		t.Error("Unexpected number of longest linked list", gbt.numOfLongestLinkedList)
	}
}

func TestInsert_NonEmpty(t *testing.T) {
	gbt := NewGameBoardHashTable()

	// create mock hash function
	hashComputer := mockHashComputer{
		stayConstant: true,
	}

	board0 := boards[0]
	_, _ = gbt.Insert(&board0, &hashComputer)

	board1 := boards[1]
	gb, ok := gbt.Insert(&board1, &hashComputer)

	if !ok {
		t.Error("Expected to insert gameboard but NOT ok")
	}

	if gb != &board1 {
		t.Error("Expected to receive", &board1, "Got", gb)
	}

	if gbt.numOfKeys != 1 {
		t.Error("Unexpected number of keys", gbt.numOfKeys)
	}

	if gbt.numOfLongestLinkedList != 2 {
		t.Error("Unexpected number of longest linked list", gbt.numOfLongestLinkedList)
	}

	linkedList := gbt.table[0]
	if linkedList.gameBoard != &board0 {
		t.Error("State of the first linked list node is broken")
	}
	if linkedList.next.gameBoard != &board1 {
		t.Error("State of the second linked list is broken")
	}
}

func TestInsert_NonEmptyLongerLinkedList(t *testing.T) {
	gbt := NewGameBoardHashTable()

	// create mock hash function
	hashComputer := mockHashComputer{
		stayConstant: true,
	}

	board0 := boards[0]
	_, _ = gbt.Insert(&board0, &hashComputer)

	board1 := boards[1]
	_, _ = gbt.Insert(&board1, &hashComputer)

	board2 := boards[2]
	gb, ok := gbt.Insert(&board2, &hashComputer)

	if !ok {
		t.Error("Expected to insert gameboard but NOT ok")
	}

	if gb != &board2 {
		t.Error("Expected to receive", &board2, "Got", gb)
	}

	if gbt.numOfKeys != 1 {
		t.Error("Unexpected number of keys", gbt.numOfKeys)
	}

	if gbt.numOfLongestLinkedList != 3 {
		t.Error("Unexpected number of longest linked list", gbt.numOfLongestLinkedList)
	}

	linkedList := gbt.table[0]
	if linkedList.gameBoard != &board0 {
		t.Error("State of the first linked list node is broken")
	}
	if linkedList.next.gameBoard != &board1 {
		t.Error("State of the second linked list is broken")
	}
	if linkedList.next.next.gameBoard != &board2 {
		t.Error("State of the third linked list is broken")
	}
}

func TestInsert_AlreadyExists(t *testing.T) {
	gbt := NewGameBoardHashTable()

	// create mock hash function
	hashComputer := mockHashComputer{
		stayConstant: true,
	}

	board0 := boards[0]
	_, _ = gbt.Insert(&board0, &hashComputer)

	board1 := boards[0]
	_, ok := gbt.Insert(&board1, &hashComputer)

	if ok {
		t.Error("Expected NOT ok Got ok")
	}

	if gbt.numOfKeys != 1 {
		t.Error("Unexpected number of keys", gbt.numOfKeys)
	}

	if gbt.numOfLongestLinkedList != 1 {
		t.Error("Unexpected number of longest linked list", gbt.numOfLongestLinkedList)
	}

	linkedList := gbt.table[0]
	if linkedList.gameBoard != &board0 {
		t.Error("State of the first linked list node is broken")
	}
	if linkedList.next != nil {
		t.Error("Expected linked list second item to be nil Got", linkedList.next)
	}
}

func TestInsert_NonEmptyNewKey(t *testing.T) {
	gbt := NewGameBoardHashTable()

	// create mock hash function
	hashComputer := mockHashComputer{}

	board0 := boards[0]
	_, _ = gbt.Insert(&board0, &hashComputer)

	board1 := boards[1]
	gb, ok := gbt.Insert(&board1, &hashComputer)

	if !ok {
		t.Error("Expected to insert gameboard but NOT ok")
	}

	if gb != &board1 {
		t.Error("Expected to receive", &board1, "Got", gb)
	}

	if gbt.numOfKeys != 2 {
		t.Error("Unexpected number of keys", gbt.numOfKeys)
	}

	if gbt.numOfLongestLinkedList != 1 {
		t.Error("Unexpected number of longest linked list", gbt.numOfLongestLinkedList)
	}
}
