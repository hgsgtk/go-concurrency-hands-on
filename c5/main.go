package main

import "fmt"

func fibonacci(c chan<- int, quit <-chan int) {
	x, y := 1, 2
	// 2.
	// mainから呼ばれたの実行する
	for {
		select {
		case c <- x:
			// 3.
			// channel受け取ったので今定義されてる x を送信する
			x, y = y, x+y
			// last おわったね
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quite := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// 1.
			// 0から始まる
			// まだ来てないブロック
			//
			// 4. 出力や!
			fmt.Println(<-c)
		}
		quite <- 0
	}()
	fibonacci(c, quite)
}
