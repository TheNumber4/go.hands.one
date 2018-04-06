package mergechannel

import (
	"reflect"
	"sync"
)

// MergeChannels merges channels doing a simple loop
func MergeChannels(channels ...chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, channel := range channels {
			for item := range channel {
				out <- item
			}
		}
	}()

	return out
}

// MergeChannelsGR merges channels using go routines
func MergeChannelsGR(channels ...chan int) chan int {
	out := make(chan int, len(channels))

	wg := new(sync.WaitGroup)
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(c chan int) {
			for item := range c {
				out <- item
			}
			wg.Done()
		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// MergeChannelsReflection merges channels using go routines
func MergeChannelsReflection(channels ...chan int) chan int {
	out := make(chan int, len(channels))

	go func() {
		defer close(out)
		// Create cases
		cases := make([]reflect.SelectCase, len(channels))
		for idx, channel := range channels {
			cases[idx] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(channel),
			}
		}

		for len(cases) != 0 {
			idx, value, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:idx], cases[idx+1:]...)
			} else {
				out <- int(value.Int())
			}
		}

	}()

	return out
}

// MergeChannelsReflectionBatched merges channels using go routines
func MergeChannelsReflectionBatched(channels ...chan int) chan []int {
	const batchSize = 100

	out := make(chan []int)

	go func() {
		defer close(out)
		// Create cases
		cases := make([]reflect.SelectCase, len(channels))
		for idx, channel := range channels {
			cases[idx] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(channel),
			}
		}

		batch := make([]int, 0, batchSize)
		for len(cases) != 0 {
			idx, value, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:idx], cases[idx+1:]...)
			} else {
				batch = append(batch, int(value.Int()))
				if len(batch) == cap(batch) {
					out <- batch
					batch = make([]int, 0, batchSize)
				}
			}
		}

		if len(batch) != 0 {
			out <- batch
		}

	}()

	return out
}
