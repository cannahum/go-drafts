// Package hashtable allows us to store all seen GameBoard objects based on some hash key
package hashtable

import (
	"github.com/cannahum/go-drafts/001-PanamaCanal/board"
)

type hashLinkedListNode struct {
	gameBoard *board.GameBoard
	next      *hashLinkedListNode
}

// HashComputer is an interface that implements the getHashKey method
type HashComputer interface {
	getHashKey(b *board.GameBoard) string
}

// GameBoardHashTable is a hash table that maps hash keys to a linked list of GameBoard items
type GameBoardHashTable struct {
	table                  map[string]*hashLinkedListNode
	numOfKeys              int
	numOfLongestLinkedList int
}

func (ht *GameBoardHashTable) has(b *board.GameBoard, getHash HashComputer) bool {
	hashKey := getHash.getHashKey(b)
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
	if !ht.has(b, getHash) {
		hashKey := getHash.getHashKey(b)
		node := ht.table[hashKey]

		if node == nil {
			node = &hashLinkedListNode{
				gameBoard: b,
			}
			ht.table[hashKey] = node
			hasInserted = true

			ht.numOfKeys++
			ht.numOfLongestLinkedList = 1
		} else {
			lengthOfThisLinkedList := 1
			for node.next != nil {
				node = node.next
				lengthOfThisLinkedList++
			}

			node.next = &hashLinkedListNode{
				gameBoard: b,
			}
			hasInserted = true
			ht.numOfLongestLinkedList = lengthOfThisLinkedList + 1
		}
	}
	return b, hasInserted
}
