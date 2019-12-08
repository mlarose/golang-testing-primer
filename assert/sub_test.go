package assert

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Basic assert example
func TestSub(t *testing.T) {
	// act
	result := Sub(4, 1)

	// assert
	assert.Equal(t, 3, result)
}

// Show the formatted message
func TestSub2(t *testing.T) {
	// act
	result := Sub(7, 2, 1)

	// assert
	assert.Equal(t, 4, result, "optional message %d != %d", result, 4)
}

// Show the wrapped test object syntax
func TestSub3(t *testing.T) {
	// act
	result := Sub(7, 9)

	// assert
	is := assert.New(t)
	is.Equal(-2, result, "optional message %d != %d", result, -2)
}

// Using test tables is also possible
func TestSubTable(tt *testing.T) {
	// arrange
	type testCases struct {
		n        []int
		expected int
	}

	tests := []testCases{
		{n: []int{0, 0}, expected: 0},
		{n: []int{2, 1}, expected: 1},
		{n: []int{1, 2}, expected: -1},
		{n: []int{4, 2, 1}, expected: 1},
	}

	for _, test := range tests {
		testName := strings.Join(strings.Fields(fmt.Sprint(test.n)), " - ")
		tt.Run(fmt.Sprintf("%s should equal %d", testName, test.expected), func(t *testing.T) {
			// act
			result := Sub(test.n...)

			// assert
			is := assert.New(t)
			is.Equal(test.expected, result)
		})
	}
}
