package basics

import (
	"fmt"
	"testing"
)

// A simple test example using the golang standard testing library
func TestAdd(t *testing.T) {
	// arrange
	args := []int{1, 2, 3}
	expected := 6

	// act
	result := Add(args...)

	// assert
	if result != expected {
		t.Errorf("result %d does not match expected sum %d", result, expected)
	}
}

// Using test cases table to reduce redundancy.
func TestAddTable(t *testing.T) {
	// arrange
	tests := []struct {
		name     string
		args     []int
		expected int
	}{
		{"2 positives", []int{2, 4}, 6},
		{"3 positives", []int{2, 4, 1}, 7},
		{"negative", []int{2, -4}, -2},
		{"zero sum", []int{1, -2, 1}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// act
			result := Add(test.args...)

			// assert
			if result != test.expected {
				t.Errorf("result %d does not match expected sum %d", result, test.expected)
			}
		})
	}
}

// Illustrating examples functioning both as unit tests and generated documentation
func ExampleAdd() {
	fmt.Println(Add(2, 3, 4))
	// output: 9
}

// Illustrating the benchmark functionality of go test
func Benchmark(b *testing.B) {
	args := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		_ = Add(append(args, i)...)
	}
}
