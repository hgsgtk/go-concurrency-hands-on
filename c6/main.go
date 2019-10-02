package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	<-time.After(10 * time.Second)
	fmt.Println("time up!")
}
