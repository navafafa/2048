package game

import (
	"math/rand"
	"fmt"
)

import "github.com/navafafa/2048/matrix"

type Game struct {
	grid matrix.Matrix
}

func (game Game) String() string {
	return game.grid.String()
}

func (game *Game) Init(i, j int) {
	game.grid.Init(i, j)
	zeros,_ := game.grid.GetZeros()
	game.grid.Set(zeros[rand.Intn(len(zeros))], rand2or4())
	zeros,_ = game.grid.GetZeros()
	game.grid.Set(zeros[rand.Intn(len(zeros))], rand2or4())
}

func (game *Game) ActOnInput(command rune) bool {
	return game.grid.Shift(command)
}

func (game *Game) GenerateNew() bool {
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

func (game *Game) Over() bool {
	_, err := game.grid.GetZeros()
	return !err
}

func rand2or4() int {
	return 2*(rand.Intn(2)+1)
}
