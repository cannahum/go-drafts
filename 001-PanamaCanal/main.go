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
			hValue = (((hValue*128 + intChar) % 100003) % 100003)
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

	i := 0
	for {
		currentBoard, err := queue.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Working loop:", i)
		i++

		if board.IsSolutionBoard(currentBoard) {
			stack.Push(currentBoard)
			doFinishSequence(&stack)
			break
		} else {
			hashTable.Insert(currentBoard, &hasher)

			variations := board.Variate(currentBoard)
			for _, v := range variations {
				if !hashTable.Has(v, &hasher) {
					queue.Enqueue(v)
				}
			}
		}
	}

	fmt.Println("Goodbye!")
}

func doFinishSequence(s *stack.GameBoardStack) {
	fmt.Println("Finish sequence")
}
