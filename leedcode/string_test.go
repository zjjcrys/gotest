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
func TestWordPattern(t *testing.T) {
	s1 := "aaaa"
	s2 := "dog cat cat dog"
	fmt.Println(wordPattern(s1, s2))
}
func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"
	fmt.Println(isAnagram(s1, s2))
}
