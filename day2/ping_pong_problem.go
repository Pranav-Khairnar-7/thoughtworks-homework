package concurrency

import (
	"fmt"
	"sync"
)

// Ping-Pong with Channels
// Problem:

// Create two goroutines:

// "Ping" sends "ping" to a channel
// "Pong" receives "ping" and responds with "pong"
// Repeat 3 times before exiting

func PingPongChanels() {

	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go ping(ch, &wg)
	go pong(ch, &wg)

	wg.Wait()

	// time.Sleep(10 * time.Second)

}

func ping(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Println("ping")
		ch <- i
		<-ch
	}
}

func pong(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		<-ch
		fmt.Println("pong")
		ch <- i
	}
}
