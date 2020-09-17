package main

import (
	"fmt"
	"bufio"
	"os"
)
import "github.com/navafafa/2048/game"

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	game := game.Game{}
	game.Init(3, 3)
	fmt.Println(game)
	for !game.Over() {
		command,_,_ := reader.ReadRune()
		if !game.ActOnInput(command) {
			fmt.Println("Can't make that move! Try again!")
			_,_,_ = reader.ReadRune()
			_,_,_ = reader.ReadRune()
			continue
		}
		_,_,_ = reader.ReadRune()
		_,_,_ = reader.ReadRune()
		game.GenerateNew()
		fmt.Println(game)
	}
	fmt.Println("Game Over!")
}
