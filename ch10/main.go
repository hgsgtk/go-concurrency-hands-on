package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
)

func main() {
	debug.SetGCPercent(-1)

	for i := 0; i < 10000; i++ {
		f := func() chan int {
			intStream := make(chan int)

			go func() {
				//defer close(intStream)
				intStream <- 3
			}()

			return intStream
		}

		ch := f()
		v := <-ch
		fmt.Println(v)
	}

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		f := func() chan int {
			intStream := make(chan int)

			go func() {
				//defer close(intStream)
				intStream <- 3
			}()

			return intStream
		}

		ch := f()
		v := <-ch
		fmt.Fprintf(w, "%v \n", v)
	})

	log.Println(http.ListenAndServe("localhost:6060", nil))
}
