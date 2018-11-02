package main

import (
	"fmt"
)

// fatal error: all goroutines are asleep - deadlock!
//
func main() {
	ch := make(chan int)
	go del(ch)

	for {
		fmt.Println(<-ch)
	}
}

func del(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
}
