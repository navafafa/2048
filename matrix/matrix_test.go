package matrix

import "testing"
import "math/rand"
//import "fmt"

var matrix Matrix = Matrix{}
const (x = 3; y = 3)

func setup(t *testing.T) {
	matrix.Init(x, y)
	if _, err := matrix.Get([2]int{x-1, y}); err {
			t.Errorf("In matrix.Init(%d, %d); for Get(%d, %d), %d is out of bounds!", x, y, x-1, y, y)
	} else if _, err := matrix.Get([2]int{x, y-1}); err {
			t.Errorf("In matrix.Init(%d, %d); for Get(%d, %d), %d is out of bounds!", x, y, x, y-1, y)
	}
	
	numOfFilled := 0
	for i:=0; i<x; i++ {
		for j:=0; j<y; j++ {
			got, err := matrix.Get([2]int{i, j})
			if got == 2 || got == 4 {
				numOfFilled++
			} else if got != 0 {
				t.Errorf("In matrix.Init(%d, %d), Get(%d, %d)=%d, want (0|2|4)", x, y, i, j, got)
			} else if !err {
				t.Errorf("In matrix.Init(%d, %d), Get(%d, %d) returned err==false, want true", x, y, i, j)
			}
		}
	}
	if numOfFilled > 0 {
		t.Errorf("In matrix.Init(%d, %d), matrix has %d non-zero entries, want exactly 0", x, y, numOfFilled)
	}
}

func setGet(t *testing.T) {
	i := rand.Intn(x)
	j := rand.Intn(y)
	v1 := 42
	v2 := 43
	matrix.Set([2]int{i, j}, v1)
	got, err := matrix.Get([2]int{i, j})
	if !err || got != v1 {
		t.Errorf("In matrix.Set([2]int{%d, %d}, %d) and matrix.Get([2]int{%d, %d}), got %d, want %d", i, j, v1, i, j, got, v1)
	}
	matrix.Set([2]int{i, j}, v2)
	matrix.Set([2]int{i, j}, v2)
	got, err = matrix.Get([2]int{i, j})
	got, err = matrix.Get([2]int{i, j})
	if !err || got != v2 {
		t.Errorf("In matrix.Set([2]int{%d, %d}, %d) and matrix.Get([2]int{%d, %d}), got %d, want %d", i, j, v2, i, j, got, v2)
	}
}

func getZeros(t *testing.T) {
	expectedZeros := make([][2]int, x*y)
	index := 0
	for i:=0; i<x; i++ {
		for j:=0; j<y; j++ {
			expectedZeros[index] = [2]int{i, j}
			index++
		}
	}
	
	matrix.Init(x, y)

	for i:=0; i<x*y; i++ {
		got, err := matrix.GetZeros()
		if err && len(expectedZeros)>0 {
			t.Errorf("In matrix.GetZeros(), no zeros left, want %d zeros", len(expectedZeros))
		} else if len(got) != len(expectedZeros) {
			t.Errorf("In matrix.GetZeros(), got %d, want %d", got, expectedZeros)
		}
		matrix.Set(expectedZeros[0], 1)
		expectedZeros = expectedZeros[1:]
	}
}

func shift(t *testing.T) {
	want := make(map[[2]int]int)
	initWant := func() {
		for i:=0; i<x; i++ {
			for j:=0; j<y; j++ {
				want[[2]int{i, j}] = 0
			}
		}
	}
	compare := func() bool {
		for i:=0; i<x; i++ {
			for j:=0; j<y; j++ {
				if v, success := matrix.Get([2]int{i, j}); !success || want[[2]int{i, j}] != v {
					return false
				}
			}
		}
		return true
	}
	dir := 'r'
	// 0 0 2      0 0 2
	// 0 0 0  ->  0 0 0
	// 0 0 2      0 0 2
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 2}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{0, 2}] = 2
	want[[2]int{2, 2}] = 2
	success := matrix.Shift(dir)
	if !compare() || success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 2 0      0 0 2
	// 0 0 0  ->  0 0 0
	// 0 2 0      0 0 2
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 1}, 2)
	matrix.Set([2]int{2, 1}, 2)
	want[[2]int{0, 2}] = 2
	want[[2]int{2, 2}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 2 0 0      0 0 2
	// 0 0 0  ->  0 0 0
	// 2 0 0      0 0 2
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 2)
	matrix.Set([2]int{2, 0}, 2)
	want[[2]int{0, 2}] = 2
	want[[2]int{2, 2}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 0
	// 0 0 0  ->  0 0 0
	// 2 0 2      0 0 4
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{2, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 0
	// 0 0 0  ->  0 0 0
	// 0 2 2      0 0 4
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{2, 1}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{2, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 0
	// 0 0 0  ->  0 0 0
	// 2 2 0      0 0 4
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{2, 1}, 2)
	want[[2]int{2, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 2 2 2      0 2 4
	// 0 0 0  ->  0 0 0
	// 2 2 2      0 2 4
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 2)
	matrix.Set([2]int{0, 1}, 2)
	matrix.Set([2]int{0, 2}, 2)
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{2, 1}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{0, 1}] = 2
	want[[2]int{0, 2}] = 4
	want[[2]int{2, 1}] = 2
	want[[2]int{2, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	dir = 'l'
	// 0 0 2      2 0 0
	// 0 0 0  ->  0 0 0
	// 0 0 2      2 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 2}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{0, 0}] = 2
	want[[2]int{2, 0}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 2 0      2 0 0
	// 0 0 0  ->  0 0 0
	// 0 2 0      2 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 1}, 2)
	matrix.Set([2]int{2, 1}, 2)
	want[[2]int{0, 0}] = 2
	want[[2]int{2, 0}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 2 0 0      2 0 0
	// 0 0 0  ->  0 0 0
	// 2 0 0      2 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 2)
	matrix.Set([2]int{2, 0}, 2)
	want[[2]int{0, 0}] = 2
	want[[2]int{2, 0}] = 2
	success = matrix.Shift(dir)
	if !compare() || success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 0
	// 0 0 0  ->  0 0 0
	// 2 0 2      4 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{2, 0}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 0
	// 0 0 0  ->  0 0 0
	// 2 2 0      4 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{2, 1}, 2)
	want[[2]int{2, 0}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 0
	// 0 0 0  ->  0 0 0
	// 0 2 2      4 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{2, 1}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{2, 0}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 2 2 2      4 2 0
	// 0 0 0  ->  0 0 0
	// 2 2 2      4 2 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 2)
	matrix.Set([2]int{0, 1}, 2)
	matrix.Set([2]int{0, 2}, 2)
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{2, 1}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{0, 0}] = 4
	want[[2]int{0, 1}] = 2
	want[[2]int{2, 0}] = 4
	want[[2]int{2, 1}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	dir = 'd'
	// 2 0 2      0	0 0
	// 0 0 0  ->  0 0 0
	// 0 0 0      2 0 2
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 2)
	matrix.Set([2]int{0, 2}, 2)
	want[[2]int{2, 0}] = 2
	want[[2]int{2, 2}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 0
	// 2 0 2  ->  0 0 0
	// 0 0 0      2 0 2
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{1, 0}, 2)
	matrix.Set([2]int{1, 2}, 2)
	want[[2]int{2, 0}] = 2
	want[[2]int{2, 2}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 0
	// 0 0 0  ->  0 0 0
	// 2 0 2      2 0 2
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{2, 0}] = 2
	want[[2]int{2, 2}] = 2
	success = matrix.Shift(dir)
	if !compare() || success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 2      0 0 0
	// 0 0 0  ->  0 0 0
	// 0 0 2      0 0 4
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 2}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{2, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 0
	// 2 0 0  ->  0 0 0
	// 2 0 0      4 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{1, 0}, 2)
	matrix.Set([2]int{2, 0}, 2)
	want[[2]int{2, 0}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 2      0 0 0
	// 0 0 2  ->  0 0 0
	// 0 0 0      0 0 4
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 2}, 2)
	matrix.Set([2]int{1, 2}, 2)
	want[[2]int{2, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 2 0 2      0 0 0
	// 2 0 2  ->  2 0 2
	// 2 0 2      4 0 4
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 2)
	matrix.Set([2]int{1, 0}, 2)
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{0, 2}, 2)
	matrix.Set([2]int{1, 2}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{1, 0}] = 2
	want[[2]int{2, 0}] = 4
	want[[2]int{1, 2}] = 2
	want[[2]int{2, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	dir = 'u'
	// 0 0 0      2	0 2
	// 0 0 0  ->  0 0 0
	// 2 0 2      0 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{0, 0}] = 2
	want[[2]int{0, 2}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      2 0 2
	// 2 0 2  ->  0 0 0
	// 0 0 0      0 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{1, 0}, 2)
	matrix.Set([2]int{1, 2}, 2)
	want[[2]int{0, 0}] = 2
	want[[2]int{0, 2}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 2 0 2      2 0 2
	// 0 0 0  ->  0 0 0
	// 0 0 0      0 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 2)
	matrix.Set([2]int{0, 2}, 2)
	want[[2]int{0, 0}] = 2
	want[[2]int{0, 2}] = 2
	success = matrix.Shift(dir)
	if !compare() || success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 2      0 0 4
	// 0 0 0  ->  0 0 0
	// 0 0 2      0 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 2}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{0, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 2 0 0      4 0 0
	// 2 0 0  ->  0 0 0
	// 0 0 0      0 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 2)
	matrix.Set([2]int{1, 0}, 2)
	want[[2]int{0, 0}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 0 0 0      0 0 4
	// 0 0 2  ->  0 0 0
	// 0 0 2      0 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{1, 2}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{0, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// 2 0 2      4 0 4
	// 2 0 2  ->  2 0 2
	// 2 0 2      0 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 2)
	matrix.Set([2]int{1, 0}, 2)
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{0, 2}, 2)
	matrix.Set([2]int{1, 2}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{0, 0}] = 4
	want[[2]int{1, 0}] = 2
	want[[2]int{0, 2}] = 4
	want[[2]int{1, 2}] = 2
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	// Filled but still possible
	// 4 4 4      4 8 4
	// 2 4 2  ->  4 2 4
	// 2 2 2      0 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{0, 0}, 4)
	matrix.Set([2]int{1, 0}, 2)
	matrix.Set([2]int{2, 0}, 2)
	matrix.Set([2]int{0, 1}, 4)
	matrix.Set([2]int{1, 1}, 4)
	matrix.Set([2]int{2, 1}, 2)
	matrix.Set([2]int{0, 2}, 4)
	matrix.Set([2]int{1, 2}, 2)
	matrix.Set([2]int{2, 2}, 2)
	want[[2]int{0, 0}] = 4
	want[[2]int{1, 0}] = 4
	want[[2]int{0, 1}] = 8
	want[[2]int{1, 1}] = 2
	want[[2]int{0, 2}] = 4
	want[[2]int{1, 2}] = 4
	success = matrix.Shift(dir)
	if !compare() || !success{
		t.Errorf("In matrix.shift(%c), got: %s \nwant: %d", dir, matrix, want)
	}

	dir='s' // invalid
	// 0 0 0      0 0 0
	// 0 2 0  ->  0 2 0
	// 0 0 0      0 0 0
	matrix.Init(x, y)
	initWant()
	matrix.Set([2]int{1, 1}, 2)
	want[[2]int{1, 1}] = 2
	success = matrix.Shift(dir)
	if success {
		t.Errorf("In matrix.shift(%c), got: true want: false", dir)
	}
}

func TestMatrix(t *testing.T){
	t.Run("Setup", func(t *testing.T){setup(t)})
	t.Run("SetGet", func(t *testing.T){setGet(t)})
	t.Run("GetZeros", func(t *testing.T){getZeros(t)})
	t.Run("shift", func(t *testing.T){shift(t)})
}