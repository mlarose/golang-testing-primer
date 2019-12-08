package basics

import (
	"fmt"
	"testing"
)

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

func ExampleAdd() {
	fmt.Println(Add(2, 3, 4))
	// output: 9
}
