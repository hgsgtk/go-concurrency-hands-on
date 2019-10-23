package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	stringStream := make(chan string)

	go func() {
		// Here!
		for _, s := range []string{"a", "b", "c"} {
			select {
			case <-done:
				return
			case stringStream <- s:
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case s := <-stringStream:
				fmt.Println(s)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	done <- struct{}{}
}
