package main

import "fmt"

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
	out := make(chan int)

	go func() {
		defer close(out)
		for number := range in {
			out <- (number * number)
		}
	}()

	return out
}

// Flatten a series of channel
func flatten(ins ...chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, in := range ins {
			for number := range in {
				out <- number
			}
		}
	}()
	return out
}

func simpleFlatten() {
	nums := []int{2, 4, 8, 16}
	fmt.Printf("Single thread: %v\n", nums)
	for nb := range flatten(
		sqr(
			emit(
				nums...))) {
		fmt.Printf("%d\n", nb)
	}
}

func workersFlatten() {
	const workerCnt = 4

	nums := []int{2, 4, 8, 16}
	fmt.Printf("2 worker thread: %v\n", nums)

	outputs := make([]chan int, workerCnt)

	input := emit(nums...)
	for i := 0; i < workerCnt; i++ {
		outputs[i] = sqr(input)
	}

	for nb := range flatten(outputs...) {
		fmt.Printf("%d\n", nb)
	}
}

func main() {
	simpleFlatten()
	workersFlatten()
}
