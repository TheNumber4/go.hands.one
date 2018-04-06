package mergechannel

import (
	"testing"
)

func buildChannels(size int) (map[int]bool, []chan int) {
	dataSet := make(map[int]bool)
	channelData := make(map[int][]int)

	for i := 0; i < size; i++ {
		dataSet[i] = true
		channelData[i%10] = append(channelData[i%10], i)
	}

	channels := make([]chan int, 0, 10)
	for _, v := range channelData {
		channels = append(channels, genIntChan(v...))
	}

	return dataSet, channels
}

func testMergeFunc(t *testing.T, f func(...chan int) chan int) {
	inputSet, channels := buildChannels(101)

	for value := range f(channels...) {
		if !inputSet[value] {
			t.Fatalf("Value %d not found !", value)
		}
		delete(inputSet, value)
	}
	if len(inputSet) != 0 {
		t.Fatalf("Not all values found !: %v", inputSet)
	}
}

func testBatchMergeFunc(t *testing.T, f func(...chan int) chan []int) {
	inputSet, channels := buildChannels(101)

	for values := range f(channels...) {
		for _, value := range values {
			if !inputSet[value] {
				t.Fatalf("Value %d not found !", value)
			}
			delete(inputSet, value)
		}
	}
	if len(inputSet) != 0 {
		t.Fatalf("Not all values found !: %v", inputSet)
	}
}

func benchMergeFunc(b *testing.B, f func(...chan int) chan int) {
	for i := 0; i < b.N; i++ {
		inputSet, channels := buildChannels(1000)

		for value := range f(channels...) {
			if !inputSet[value] {
				b.Fatalf("Value %d not found !", value)
			}
			delete(inputSet, value)
		}
		if len(inputSet) != 0 {
			b.Fatalf("Not all values found !: %v", inputSet)
		}
	}
}

func benchBatchMergeFunc(b *testing.B, f func(...chan int) chan []int) {
	for i := 0; i < b.N; i++ {
		inputSet, channels := buildChannels(1000)

		for values := range f(channels...) {
			for _, value := range values {
				if !inputSet[value] {
					b.Fatalf("Value %d not found !", value)
				}
				delete(inputSet, value)
			}
		}
		if len(inputSet) != 0 {
			b.Fatalf("Not all values found !: %v", inputSet)
		}
	}
}

func TestMergeChannels(t *testing.T) {
	testMergeFunc(t, MergeChannels)
}

func TestMergeChannelsGR(t *testing.T) {
	testMergeFunc(t, MergeChannelsGR)
}

func TestMergeChannelsReflection(t *testing.T) {
	testMergeFunc(t, MergeChannelsReflection)
}

func TestMergeChannelsReflectionBatched(t *testing.T) {
	testBatchMergeFunc(t, MergeChannelsReflectionBatched)
}

func BenchmarkMergeChannels(b *testing.B) {
	benchMergeFunc(b, MergeChannels)
}

func BenchmarkMergeChannelsGR(b *testing.B) {
	benchMergeFunc(b, MergeChannelsGR)
}

func BenchmarkMergeChannelsReflection(b *testing.B) {
	benchMergeFunc(b, MergeChannelsReflection)
}

func BenchmarkMergeChannelsReflectionBatched(b *testing.B) {
	benchBatchMergeFunc(b, MergeChannelsReflectionBatched)
}
