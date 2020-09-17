package game

import "testing"
import "math/rand"
import "fmt"

var game Game = Game{}
const (x=3; y=3)

func setup(t *testing.T) {
	game.Init(x, y)
	
	if got, _ := game.grid.GetZeros(); len(got) != x*y-2 {
		t.Errorf("In Init() for the number of non zero elements got: %d want: 2", len(got))
	}
}

func generateNew(t *testing.T) {
	game.Init(x, y)

	for err := false; !err; err = game.GenerateNew() {
		game.ActOnInput('r')
	}
	if got, err := game.grid.GetZeros(); !err {
		t.Errorf("In GenerateNew() the function returned while there are still %d zeros left", len(got))
	}
}

func over(t *testing.T) {
	for game.Init(x, y); !game.Over(); game.GenerateNew() {
		fmt.Println(game)
		moves := [4]rune{'r', 'l', 'u', 'd'}
		move := moves[rand.Intn(len(moves))]
		game.ActOnInput(move)
		fmt.Println(string(move))
	}
	fmt.Println(game)
	if game.ActOnInput('r') == true || game.ActOnInput('l') == true || game.ActOnInput('u') == true || game.ActOnInput('d') == true {
		t.Errorf("In Over() the function returned true while there are still possible moves")
	}
}

func TestGame(t *testing.T) {
	t.Run("Setup", func(t *testing.T){setup(t)})
	t.Run("GenerateNew", func(t *testing.T){generateNew(t)})
	t.Run("Over", func(t *testing.T){over(t)})
}