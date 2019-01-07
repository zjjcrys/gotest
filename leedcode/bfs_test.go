package leedcode

import (
	"fmt"
	"testing"
)

func TestLadderLength(t *testing.T) {
	begin := "hit"
	end := "cog"
	list := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(begin, end, list))
}
