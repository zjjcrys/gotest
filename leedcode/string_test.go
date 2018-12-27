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
func TestIsScramble(t *testing.T) {
	s1 := "ccabcbabcbabbbbcbb"
	s2 := "bbbbabccccbbbabcba"
	fmt.Println(isScramble(s1, s2))
}
