package main

import "fmt"

func main() {
	origins := make(chan int)

	go counter(origins)

	calculated := make(chan int)

	go doubler(calculated, origins)

	printer(calculated)
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func doubler(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * 2
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
