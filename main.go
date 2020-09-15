package main

import (
	"fmt"
)

import "github.com/navafafa/2048/game"

func main() {
	game := game.Game{}
	game.Init(3, 3)
	fmt.Println(game.grid)
	for !game.Over() {
		game.ActOnInput()
		_,_,_ = game.reader.ReadRune()
		_,_,_ = game.reader.ReadRune()
		game.GenerateNew()
		fmt.Println(game.grid)
	}
}
