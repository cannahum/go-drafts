// Package hashtable allows us to store all seen GameBoard objects based on some hash key
package hashtable

import (
	"github.com/cannahum/go-drafts/001-PanamaCanal/board"
)

type hashLinkedListNode struct {
	gameBoard *board.GameBoard
	next      *hashLinkedListNode
}

// HashComputer is an interface that implements the GetHashKey method
type HashComputer interface {
	GetHashKey(b *board.GameBoard) int
}

// GameBoardHashTable is the object in which we keep all the "seen" boards
type GameBoardHashTable struct {
	table                  map[int]*hashLinkedListNode
	numOfKeys              int
	numOfLongestLinkedList int
}

// GetNumOfKeys returns the number of keys on the hash table.
func (ht *GameBoardHashTable) GetNumOfKeys() int {
	return ht.numOfKeys
}

// GetNumOfLongestLinkedList returns the number of keys on the hash table.
func (ht *GameBoardHashTable) GetNumOfLongestLinkedList() int {
	return ht.numOfLongestLinkedList
}

// NewGameBoardHashTable Constructs a GameBoardHashTable object
func NewGameBoardHashTable() GameBoardHashTable {
	return GameBoardHashTable{
		table:                  map[int]*hashLinkedListNode{},
		numOfKeys:              0,
		numOfLongestLinkedList: 0,
	}
}

// Has checks if a given board is already "seen"; in other words, it exists in the table
func (ht *GameBoardHashTable) Has(b *board.GameBoard, getHash HashComputer) bool {
	hashKey := getHash.GetHashKey(b)
	node := ht.table[hashKey]

	for node != nil {
		gb := node.gameBoard
		if gb.Board == b.Board {
			return true
		}
		node = node.next
	}

	return false
}

// Insert checks if the hashtable already has this board or not. If not,
// it adds it
func (ht *GameBoardHashTable) Insert(b *board.GameBoard, getHash HashComputer) (*board.GameBoard, bool) {
	hasInserted := false

	if !ht.Has(b, getHash) {
		hashKey := getHash.GetHashKey(b)
		node := ht.table[hashKey]
		var lengthOfThisLinkedList int

		if node == nil {
			node = &hashLinkedListNode{
				gameBoard: b,
			}
			ht.table[hashKey] = node
			hasInserted = true
			ht.numOfKeys++
		} else {
			lengthOfThisLinkedList = 1
			for node.next != nil {
				node = node.next
				lengthOfThisLinkedList++
			}

			node.next = &hashLinkedListNode{
				gameBoard: b,
			}
			hasInserted = true
			lengthOfThisLinkedList++
		}

		if lengthOfThisLinkedList > ht.numOfLongestLinkedList {
			ht.numOfLongestLinkedList = lengthOfThisLinkedList
		}
	}
	return b, hasInserted
}
