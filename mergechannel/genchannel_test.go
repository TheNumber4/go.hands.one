package mergechannel

import "testing"

func TestGenIntChan(t *testing.T) {
	input := []int{1, 2, 3, 4}
	c := genIntChan(input...)

	i := 1
	for data := range c {
		if data != i {
			t.Fatalf("Expected %d but got %d", i, data)
		}
		i++
	}
}
