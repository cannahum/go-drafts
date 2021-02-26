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

	fmt.Println(gameBoard.GetCurrentBoardState())
	players := []player.Player{playerX, playerO}
	playerIndex := 0
	for {
		p, _ := players[playerIndex].(player.Player)
		gameFinished, err := gameBoard.RegisterMove(*p.MakeAMove(0, 0))
		if err != nil {
			panic("Error")
		}
		if gameFinished {
			fmt.Println("Game finished!")
			return
		}
		playerIndex = (playerIndex + 1) % 2
	}
}
