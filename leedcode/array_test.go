package leedcode

import (
	"fmt"
	"testing"
)

/*
test func:go test -v ./libraries/im/*.go -test.run TestSendPush
*/
func TestSortArrayByParity(t *testing.T) {
	s := make([]int, 4)
	s[0] = 3
	s[1] = 1
	s[2] = 2
	s[3] = 4
	fmt.Println(SortArrayByParity(s))
}
func TestTranspose(t *testing.T) {
	twoD := make([][]int, 2)
	for i := 0; i < 2; i++ {
		innerLen := 3
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(Transpose(twoD))
}

func TestCombinationSum3(t *testing.T) {
	fmt.Println(CombinationSum3(3, 9))
}

func TestGameOfLife(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	GameOfLife(board)
	fmt.Println(board)
}
func TestMajorityElement(t *testing.T) {
	board := []int{}
	fmt.Println(MajorityElement(board))
}
