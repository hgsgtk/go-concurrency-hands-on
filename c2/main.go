package main

import (
	"fmt"
	"time"
)

func main() {
	gn := 1
	go func() {
		for i := 2; i <= 5; i++ {
			fmt.Printf("%d * %d\n", gn, i)
			gn *= i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d + %d\n", gn, i)
		gn += i
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(gn)
}
