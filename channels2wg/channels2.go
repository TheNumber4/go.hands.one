package main

import (
	"fmt"
	"sync"
)

// emit emits a serie of value as a channel
func emit(numbers ...int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, number := range numbers {
			out <- number
		}
	}()

	return out
}

// sqr will sqr root a number
func sqr(in chan int) chan int {
	// Create a wait group
	wg := new(sync.WaitGroup)

	out := make(chan int)

	// Create the workers - they will all output to the same out channel
	const workerCnt = 4
	for i := 0; i < workerCnt; i++ {
		// Add to the worker count
		wg.Add(1)
		// Create a goroutine closure
		go func() {
			for number := range in {
				out <- (number * number)
			}
			// we are done
			wg.Done()
		}()
	}

	// Monitor the wait group in another gorouting to close the channel
	go func() {
		wg.Wait()
		close(out)
	}()

	// Return output chanel
	return out
}

func main() {
	nums := []int{2, 4, 8, 16}
	fmt.Printf("WG thread: %v\n", nums)
	for nb := range sqr(
		emit(nums...)) {
		fmt.Printf("%d\n", nb)
	}
}
