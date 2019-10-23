package main

import "fmt"

func main() {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		intStream <- 3
	}()

	// 受け取れるよ
	integer, ok := <-intStream
	fmt.Printf("(%v) %v", ok, integer)
}
