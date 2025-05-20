package concurrency

import (
	"fmt"
	"sync"
)

// Concurrent Sum Calculator
// Problem:
// Write a program that:
// Splits a slice of numbers into two halves
// Calculates the sum of each half in separate goroutines
// Combines the results and prints the total sum

func ConcurrentSumWithChanels() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var answer int
	ch := make(chan int, 2)

	go sumArr(arr[:5], ch)
	go sumArr(arr[5:], ch)

	for i := 0; i < 2; i++ {
		answer += <-ch
	}

	fmt.Println("Sum ==>", answer)

}

func sumArr(arr []int, ch chan int) {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	ch <- sum
}

func sumArrWithWg(arr []int, sum *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < len(arr); i++ {
		*sum += arr[i]
	}
}

func ConcurrentSumWithWg() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var sum1 int
	var sum2 int

	wg := sync.WaitGroup{}

	wg.Add(2)

	go sumArrWithWg(arr[:5], &sum1, &wg)
	go sumArrWithWg(arr[5:], &sum2, &wg)

	wg.Wait()

	fmt.Println("Sum ==>", sum1+sum2)

}
