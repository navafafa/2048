package game

import (
	"math/rand"
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
	if err {
		return true
	}
	game.grid.Set(zeros[rand.Intn(len(zeros))], rand2or4())
	return false
}

func (game *Game) Over() bool {
	if _, noZeros := game.grid.GetZeros(); !noZeros {
		return false
	}
	
	for i:=0; i<game.grid.GetX(); i++ {
		for j:=0; j<game.grid.GetY(); j++ {
			if i != game.grid.GetX()-1 {
				v1,_ := game.grid.Get([2]int{i, j})
				v2,_ := game.grid.Get([2]int{i+1, j})
				if v1 == v2 {
					return false
				}
			}
			if j != game.grid.GetY()-1 {
				v1,_ := game.grid.Get([2]int{i, j})
				v2,_ := game.grid.Get([2]int{i, j+1})
				if v1 == v2 {
					return false
				}
			}
		}
	}
	
	return true
}

func rand2or4() int {
	return 2*(rand.Intn(2)+1)
}
