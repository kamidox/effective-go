package main

import (
	"fmt"
	"time"
)

func handler(q chan int) {
	fmt.Println("start handler ...")
	for r := range q {
		fmt.Println("handle request ...", r)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("done", r)
	}
}

func serve(q chan int, done chan int) {
	// 创建了 10 个 goroutine 来并行处理请求，最多只能同时处理 10 个请求
	for i := 0; i < 10; i++ {
		go handler(q)
	}
	<-done
}

func mockRequest(q chan int, done chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("start to send mock requests ...")
	// 模拟 20 个请求几乎同时到达，可以看到，最早到达的 10 个请求在并行处理。此后，如果有请求处理完，才会有新的请求可以处理。
	for i := 0; i < 20; i++ {
		q <- i
	}
	time.Sleep(1 * time.Second)
	fmt.Println("all mock requests done.")
	done <- 1
}

func main() {
	q := make(chan int, 20)
	done := make(chan int)
	go mockRequest(q, done)
	serve(q, done)
}
