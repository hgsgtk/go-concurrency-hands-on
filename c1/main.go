package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main goroutine start")
	defer fmt.Println("main goroutine done")

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("goroutine 1 done")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine 2 done")
	}()

	time.Sleep(5 * time.Second)
}
