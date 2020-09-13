package main

import (
	"fmt"
	"math/rand"
	"bufio"
	"os"
)

import "github.com/navafafa/2048"

type Game struct {
	grid matrix.Matrix
	reader	 *bufio.Reader
}

func (game *Game) init(i, j int) {
	game.grid.Init(i, j)
	zeros,_ := game.grid.GetZeros()
	game.grid.Set(zeros[rand.Intn(len(zeros))], rand2or4())
	zeros,_ = game.grid.GetZeros()
	game.grid.Set(zeros[rand.Intn(len(zeros))], rand2or4())
	game.reader = bufio.NewReader(os.Stdin)
}

func (game *Game) actOnInput() {
	command,_,_ := game.reader.ReadRune()
	for !game.grid.Shift(command) {
		fmt.Println("Can't make that move! Try again!")
		_,_,_ = game.reader.ReadRune()
		_,_,_ = game.reader.ReadRune()
		command,_,_ = game.reader.ReadRune()
	}
}

func (game *Game) generateNew() bool {
	zeros, err := game.grid.GetZeros()
	if !err {
		return err
	}
	game.grid.Set(zeros[rand.Intn(len(zeros))], rand2or4())
	zeros,_ = game.grid.GetZeros()
	fmt.Println(zeros)
	if err {
		game.grid.Set(zeros[rand.Intn(len(zeros))], rand2or4())
	}
	return true
}

func (game *Game) over() bool {
	_, err := game.grid.GetZeros()
	return !err
}

func rand2or4() int {
	return 2*(rand.Intn(2)+1)
}

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
