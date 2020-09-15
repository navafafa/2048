package main

import (
	"fmt"
)


func main() {
	game := Game{}
	game.init(3, 3)
	fmt.Println(game.grid)
	for !game.over() {
		game.actOnInput()
		_,_,_ = game.reader.ReadRune()
		_,_,_ = game.reader.ReadRune()
		game.generateNew()
		fmt.Println(game.grid)
	}
}
