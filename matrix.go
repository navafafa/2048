package matrix

import (
	"fmt"
	"strconv"
	"math"
)

type Matrix struct {
	i, j int
	data map[[2]int]int
}

func (matrix *Matrix) Init(i, j int) {
	matrix.i = i
	matrix.j = j
	matrix.data = make(map[[2]int]int)
	for x:=0; x<matrix.i; x++ {
		for y:=0; y<matrix.j; y++ {
			matrix.data[[2]int{x, y}] = 0
		}
	}
}

func (matrix Matrix) String() string{
	str := ""
	for i:=0; i<matrix.i; i++ {
		str += fmt.Sprintf("\n")
		str += fmt.Sprintf("| ")
		for j:=0; j<matrix.j; j++ {
			str += fmt.Sprintf(strconv.Itoa(matrix.data[[2]int{i, j}]))
			str += fmt.Sprintf(" | ")
		}
	}
	return str
}

func (matrix *Matrix) Shift(direction rune) bool{
	type LoopParams struct {
		low, high, step int
	}
	paramsA := LoopParams{}
	paramsB := LoopParams{}
	straight := true
	err := false
	
	switch direction {
		case 'r':
			paramsA.low = 0
			paramsA.high = matrix.i
			paramsA.step = 1
			paramsB.low = matrix.j-1
			paramsB.high = 0
			paramsB.step = -1
		case 'l':
			paramsA.low = 0
			paramsA.high = matrix.i
			paramsA.step = 1
			paramsB.low = 0
			paramsB.high = matrix.j-1
			paramsB.step = 1
		case 'd':
			paramsA.low = 0
			paramsA.high = matrix.j
			paramsA.step = 1
			paramsB.low = matrix.i-1
			paramsB.high = 0
			paramsB.step = -1
			straight = false
		case 'u':
			paramsA.low = 0
			paramsA.high = matrix.j
			paramsA.step = 1
			paramsB.low = 0
			paramsB.high = matrix.i-1
			paramsB.step = 1
			straight = false
		default:
			return err
	}

	index := func(i, j int) [2]int {
		if straight {
			return [2]int{i, j}
		}
		return [2]int{j, i}
	}

	for i:=paramsA.low; i!=paramsA.high; i+=paramsA.step {
		for j:=paramsB.low; j!=paramsB.high; j+=paramsB.step {
			if matrix.data[index(i, j)] == matrix.data[index(i, j+paramsB.step)] {
				matrix.data[index(i, j)] = matrix.data[index(i, j)] + matrix.data[index(i, j+paramsB.step)]
				matrix.data[index(i, j+paramsB.step)] = 0
			}
		}
	}
	for round:=0; round<int(math.Abs(float64(paramsB.high-paramsB.low))); round++ {
		for i:=paramsA.low; i!=paramsA.high; i+=paramsA.step {
			for j:=paramsB.high; j!=paramsB.low; j-=paramsB.step {
				if matrix.data[index(i, j-paramsB.step)] == 0{
					matrix.data[index(i, j-paramsB.step)] = matrix.data[index(i, j)]
					matrix.data[index(i, j)] = 0
					err = true
				}
			}
		}
	}
	return err
}

func (matrix *Matrix) Get(coordinates [2]int) (int, bool) {
	v, e := matrix.data[coordinates]
	return v, e
}

func (matrix *Matrix) Set(coordinates [2]int, v int) bool {
	if _,err := matrix.get(coordinates); !err {
		return err
	}
	matrix.data[coordinates] = v
	return true
}

func (matrix *Matrix) GetZeros() ([][2]int, bool) {
	err := false
	zeros := [][2]int{}
	for k, v := range matrix.data {
		if v == 0 {
			zeros = append(zeros, k)
			err = true
		}
	}
	return zeros, err
}
