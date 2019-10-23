package main

import (
	"fmt"
	"sync"
	"time"
)

func zikankakaru(wg *sync.WaitGroup, s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println(s)

	wg.Done()
}

func main() {
	start := time.Now()

	wg := &sync.WaitGroup{}

	fmt.Println("hello")
	for i := 0; i < 10; i++ {
		wg.Add(1)

		zikankakaru(wg, "world")
	}

	wg.Wait()

	end := time.Now()
	fmt.Printf("処理時間: %f 秒\n", (end.Sub(start)).Seconds())
}
