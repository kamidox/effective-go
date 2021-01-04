package main

import (
	"fmt"
)

func ping(c chan int, v int) {
	fmt.Println("ping ...", v)
	c <- v
}

func pong(c chan int, done chan int) {
	v := <-c
	fmt.Println("pong ...", v)
	done <- v

}

func main() {
	c := make(chan int)
	done := make(chan int)
	fmt.Println("running ping pong with goroutine")
	for i := 10; i >= 0; i-- {
		go ping(c, i)
		go pong(c, done)
	}
	for i := 10; i >= 0; i-- {
		<-done
	}
	fmt.Println("done")
}
