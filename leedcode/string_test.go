package leedcode

import (
	"fmt"
	"testing"
)

func TestFullJustify(t *testing.T) {
	board := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	fmt.Println(fullJustify(board, 16))
}
func TestMinWindow(t *testing.T) {
	sStr := "ADOBECODEBANC"
	tStr := "ABC"
	fmt.Println(minWindow(sStr, tStr))
}
