package main

import (
	"fmt"

	"github.com/pijalu/go.hands.one/slice/queue"
)

func main() {
	var q queue.Queue

	q.Add(1, 2, 3)

	for {
		res := q.Next()
		if res == nil {
			fmt.Printf("Queue is empty\n")
			break
		}
		fmt.Printf("Got %d\n", res.(int))
	}
}
