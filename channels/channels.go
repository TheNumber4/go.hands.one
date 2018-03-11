package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type wordCount map[string]int

// emit Read file line by line and output as a channel
func emit(filename string) chan wordCount {
	out := make(chan wordCount)
	// Within a goroutine
	go func() {
		// close the channel we created when we are done
		defer close(out)

		// Open source file
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		// Close when done
		defer f.Close()

		// We are going to read it line by line
		scanner := bufio.NewScanner(f)

		// Read a line, split by word and emit the word with count 1
		for scanner.Scan() {
			text := strings.ToLower(scanner.Text())
			for _, word := range strings.Fields(text) {
				m := make(wordCount)
				m[word] = 1
				out <- m
			}
		}
	}()

	// return the channel already
	return out
}

// reduce takes a series of word count channels and merge them to a single channel
func reduce(ins ...chan wordCount) chan wordCount {
	out := make(chan wordCount)
	// with a gorouting
	go func() {
		// close the result channel when work is done
		defer close(out)
		// Our merge result
		m := make(map[string]int)
		for _, in := range ins {
			for item := range in {
				for k, v := range item {
					m[k] += v
				}
			}
		}
		// Push the merge
		out <- m
	}()
	return out
}

// Number of reduce go routine
const NB_WORKER = 4

func main() {
	// Some usage
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "USAGE: %s file...\n", os.Args[0])
		os.Exit(42)
	}

	// Create dedicated input channel for each file
	inputs := make([]chan wordCount, len(os.Args)-1)
	for idx, file := range os.Args[1:] {
		inputs[idx] = emit(file)
	}

	// Create output for each of the worker
	// they will all work on the series of output channels
	outputs := make([]chan wordCount, NB_WORKER)
	for i := 0; i < NB_WORKER; i++ {
		outputs[i] = reduce(inputs...)
	}

	// Reduce the NB_WORKER outputs to a single map
	for output := range reduce(outputs...) {
		for k, v := range output {
			fmt.Printf("%s = %d\n", k, v)
		}
	}
}
