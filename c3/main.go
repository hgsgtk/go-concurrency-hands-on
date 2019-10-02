package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 2)
		done <- struct{}{}
	}()
	<-done
	fmt.Println("done")
}
