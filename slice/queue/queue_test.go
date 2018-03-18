package queue

import (
	"fmt"
	"testing"
) // github.com/pijalu/go.hands.one/slice/queue

func TestAddNextInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	q := new(Queue)
	q.Add(1, 2, 3, 4, 5)

	for _, expected := range input {
		actual := q.Next().(int)
		if expected != actual {
			t.Fatalf("Expected %d but got %d", expected, actual)
		}
	}
}

func TestAddNextString(t *testing.T) {
	input := []string{"one", "two", "three", "four", "five"}

	q := new(Queue)
	q.Add("one", "two", "three", "four", "five")

	for _, expected := range input {
		actual := q.Next().(string)
		if expected != actual {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	}
}

func TestSizeInt(t *testing.T) {
	q := new(Queue)
	q.Add(1, 2, 3, 4, 5)
	expectedSize := 5
	actualSize := q.Size()
	if actualSize != expectedSize {
		t.Fatalf("Expected %d elements but got %d", expectedSize, actualSize)
	}
}

func TestClearInt(t *testing.T) {
	q := new(Queue)
	q.Add(1, 2, 3, 4, 5)
	q.Clear()

	expectedSize := 0
	actualSize := q.Size()
	if actualSize != expectedSize {
		t.Fatalf("Expected %d elements but got %d", expectedSize, actualSize)
	}
}

func TestStringer(t *testing.T) {
	q := new(Queue)
	q.Add(1, 2)
	expected := "Q:[1 2]"

	// Sprint should call Stringer method
	if fmt.Sprint(q) != expected {
		t.Fatalf("Expected %s but got %s", expected, fmt.Sprint(q))
	}
}
