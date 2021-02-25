package main

import (
	"fmt"

	"github.com/cannahum/go-drafts/003-TicTacToe/gameboard"
	"github.com/cannahum/go-drafts/003-TicTacToe/move"
	"github.com/cannahum/go-drafts/003-TicTacToe/player"
)

func main() {
	fmt.Println("Hey!")

	gameBoard := gameboard.NewGameBoard()
	playerX := player.NewPlayer(move.X)
	playerO := player.NewPlayer(move.O)
	fmt.Println(gameBoard)
}
