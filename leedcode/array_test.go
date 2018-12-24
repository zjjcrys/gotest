package leedcode

import (
	"fmt"
	"testing"
)

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
func TestSolveSudoku(t *testing.T) {
	board := [][]byte{
		{'5', 3, '.', '.', 7, '.', '.', '.', '.'},
		{6, '.', '.', 1, 9, 5, '.', '.', '.'},
		{'.', 9, 8, '.', '.', '.', '.', 6, '.'},
		{8, '.', '.', '.', 6, '.', '.', '.', 3},
		{4, '.', '.', 8, '.', 3, '.', '.', 1},
		{7, '.', '.', '.', 2, '.', '.', '.', 6},
		{'.', 6, '.', '.', '.', '.', 2, 8, '.'},
		{'.', '.', '.', 4, 1, 9, '.', '.', 5},
		{'.', '.', '.', '.', 8, '.', '.', 7, 9}}
	solveSudoku(board)
}
func TestIsMatch(t *testing.T) {
	s := "ho"
	p := "ho***"
	fmt.Println(isMatch2(s, p))
}
