package main

import (
	"fmt"

	"github.com/cannahum/go-drafts/003-TicTacToe/gameboard"
)

func main() {
	fmt.Println("Hey!")

	gameBoard := gameboard.NewGameBoard()
	// playerX := player.NewPlayer(move.X)
	// playerO := player.NewPlayer(move.O)

	fmt.Println(gameBoard.GetCurrentBoardState())
}
