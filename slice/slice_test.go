package main

import (
	"fmt"
	"testing"
)

func printSliceInfo(text string, slice []string) {
	fmt.Printf("%s = %v (len: %d - cap: %d)\n",
		text,
		slice,
		len(slice),
		cap(slice))
}

func TestSlice(t *testing.T) {
	a := []string{"a", "b", "c"}

	type testData struct {
		name  string
		slice []string
	}

	for _, slice := range []testData{
		{name: "a[]", slice: a},
		{name: "a[1:]", slice: a[1:]},
		{name: "a[:1]", slice: a[:1]},
	} {
		printSliceInfo(slice.name, slice.slice)
	}

}
