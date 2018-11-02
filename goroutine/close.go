package main

import (
	"fmt"
)

// fatal error: all goroutines are asleep - deadlock! 因为一个协程已经执行完，但是main线程等待读取数据，永远等不到
// close.go and select.go solve the question
func main() {
	ch := make(chan int)
	go start(ch)

	for {
		if i, ok := <-ch; ok {
			fmt.Println(i)
		}
	}
}

func start(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}
