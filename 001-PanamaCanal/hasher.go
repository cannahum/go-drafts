package main

import "github.com/cannahum/go-drafts/001-PanamaCanal/board"

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
