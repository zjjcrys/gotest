package leedcode

import (
	"fmt"
	"testing"
)

func TestNumberOfArithmeticSlices(t *testing.T) {
	board := []int{1, 2, 3, 4}
	fmt.Println(NumberOfArithmeticSlices(board))
}
func TestLengthOfLIS(t *testing.T) {
	board := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(board))
}
