package main

import (
	"fmt"
	"log"

	"github.com/cannahum/go-drafts/001-PanamaCanal/board"
	"github.com/cannahum/go-drafts/001-PanamaCanal/hashtable"
	"github.com/cannahum/go-drafts/001-PanamaCanal/queue"
	"github.com/cannahum/go-drafts/001-PanamaCanal/stack"
)

func main() {
	fmt.Println("Welcome to the Panama Canal Puzzle")

	q := queue.GameBoardQueue{}
	s := stack.GameBoardStack{}
	hashTable := hashtable.NewGameBoardHashTable()
	hasher := hashComputer{}

	firstBoard := board.GameBoard{
		Board: board.Board{
			{"C", "A", "N", "A", "M", "A"},
			{"P", "A", "N", "A", "L", ""},
		},
	}

	q.Enqueue(&firstBoard)

	for {
		currentBoard, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		if !hashTable.Has(currentBoard, &hasher) {
			if board.IsSolutionBoard(currentBoard) {
				prepareStack(currentBoard, &s)
				doFinishSequence(&s)
				break
			} else {
				hashTable.Insert(currentBoard, &hasher)

				variations := board.Variate(currentBoard)
				for _, v := range variations {
					q.Enqueue(v)
				}
			}
		}
	}

	// Print statistics
	fmt.Println("Data Structures and Stats")
	fmt.Printf("HashTable: # of Keys: %d, Key with largest Linked List has %d nodes\n", hashTable.GetNumOfKeys(), hashTable.GetNumOfLongestLinkedList())
	fmt.Printf("Queue: Current size: %d, The maximum it ever got was: %d\n", q.GetCurrentLength(), q.GetMaxLength())
	fmt.Printf("Stack: Current size: %d, The maximum it ever got was: %d\n", s.GetCurrentLength(), s.GetMaxLength())
	fmt.Println("Goodbye!")
}

func prepareStack(b *board.GameBoard, s *stack.GameBoardStack) {
	for b != nil {
		s.Push(b)
		b = b.GetPrev()
	}
}

func doFinishSequence(s *stack.GameBoardStack) {
	fmt.Println("Found a solution!")

	i := 0
	for {
		b, err := s.Pop()

		if err != nil {
			break
		}

		if i == 0 {
			fmt.Printf("Step: %d:\n", i)
		} else {
			fmt.Printf("Step: %d: Move %s %s\n", i, b.GetChar(), b.GetDirection())
		}
		fmt.Println(b)
		fmt.Println()
		i++
	}
}
