package main

import (
	"fmt"
	"github.com/cannahum/go-drafts/001-PanamaCanal/board"
	"github.com/cannahum/go-drafts/001-PanamaCanal/hashtable"
	"github.com/cannahum/go-drafts/001-PanamaCanal/queue"
	"github.com/cannahum/go-drafts/001-PanamaCanal/stack"
	"log"
)

type hashComputer struct{}

func (h *hashComputer) GetHashKey(b *board.GameBoard) int {
	hValue := 0
	for _, line := range b.Board {
		for _, char := range line {
			if char == "" {
				char = " "
			}
			intChar := int(char[0])
			hValue = ((hValue*128 + intChar) % 100003) % 100003
		}
	}
	return hValue
}

func main() {
	fmt.Println("Welcome to the Panama Canal Puzzle")

	queue := queue.GameBoardQueue{}
	stack := stack.GameBoardStack{}
	hashTable := hashtable.NewGameBoardHashTable()
	hasher := hashComputer{}

	firstBoard := board.GameBoard{
		Board: board.Board{
			{"C", "A", "N", "A", "M", "A"},
			{"P", "A", "N", "A", "L", ""},
		},
	}

	queue.Enqueue(&firstBoard)

	for {
		currentBoard, err := queue.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		if !hashTable.Has(currentBoard, &hasher) {
			if board.IsSolutionBoard(currentBoard) {
				prepareStack(currentBoard, &stack)
				doFinishSequence(&stack)
				break
			} else {
				hashTable.Insert(currentBoard, &hasher)

				variations := board.Variate(currentBoard)
				for _, v := range variations {
					queue.Enqueue(v)
				}
			}
		}
	}

	// Print statistics
	fmt.Println("Data Structures and Stats")
	fmt.Printf("HashTable: # of Keys: %d, Key with largest Linked List has %d nodes\n", hashTable.GetNumOfKeys(), hashTable.GetNumOfLongestLinkedList())
	fmt.Printf("Queue: Current size: %d, The maximum it ever got was: %d\n", queue.GetCurrentLength(), queue.GetMaxLength())
	fmt.Printf("Stack: Current size: %d, The maximum it ever got was: %d\n", stack.GetCurrentLength(), stack.GetMaxLength())
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
