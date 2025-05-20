package concurrency

import (
	"fmt"
	"pranav/util"
	"sync"
	"time"
)

// problem 1
// Create a program that launches 5 goroutines. Each goroutine should:

// Take a name as input (e.g., "Alice", "Bob")
// Sleep for a random time (1-3 seconds)
// Print "Hello, [name]!"

func ConcurrentGreeter() {
	fmt.Println("ConcurrentGreeter called")
	var wg sync.WaitGroup

	wg.Add(5)

	go greet("Alice", &wg)
	go greet("BOB", &wg)
	go greet("Pranav", &wg)
	go greet("Prateek", &wg)
	go greet("Vinay", &wg)

	wg.Wait()

}

func greet(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	num := util.GetRandomUpto(2)
	fmt.Printf("Sleep time for %s is %d \n", name, num)
	time.Sleep(time.Duration(num) * time.Second)
	fmt.Printf("Hello %s!!! \n", name)
}
