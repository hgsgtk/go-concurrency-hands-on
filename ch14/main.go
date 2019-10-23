package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := new(sync.Mutex) // add
	c := make(chan bool)

	for i := 0; i < 5; i++ {
		m.Lock() // add
		go func(i int) {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(i)

			m.Unlock() // add

			c <- true
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
	}
}
