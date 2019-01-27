package maxspan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSpan(t *testing.T) {
	// TEST: integer slice
	inputs := [][]int{
		[]int{1},
		[]int{1, 2, 1, 1, 3},
		[]int{1, 4, 2, 1, 4, 1, 4},
		[]int{1, 4, 2, 1, 4, 4, 4},
	}
	want := []int{1, 4, 6, 6}
	for index, input := range inputs {
		assert.Equal(t, want[index], maxSpanInt(input), "Max span between to integer slice elements should be returned")
	}

	// TEST: slice of any type
	inputsInterface := [][]interface{}{
		[]interface{}{1, 4, 2, 1, 4, 1, 4},
		[]interface{}{"a", "b", "c", "d", "a", "c"},
		[]interface{}{true, true, true, false, false, true},
	}
	want = []int{6, 5, 6}
	for index, input := range inputsInterface {
		assert.Equal(t, want[index], maxSpan(input), "Max span between to slice elements should be returned")
	}
}
