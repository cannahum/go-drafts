package main

import (
	"fmt"

	"github.com/cannahum/go-drafts/003-TicTacToe/gameboard"
)

func main() {
	fmt.Println("Hey!")

	gameBoard := gameboard.NewGameBoard()
	fmt.Println(gameBoard)
}
