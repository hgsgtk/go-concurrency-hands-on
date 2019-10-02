package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		fmt.Printf("send %d\n", i)
		ch <- i
	}

	for i := 0; i < 5; i++ {
		go func() {
			n := <-ch
			fmt.Printf("receive %d\n", n)
		}()
	}
	time.Sleep(3 * time.Second)
}
