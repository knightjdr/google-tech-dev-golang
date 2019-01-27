package maxspan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSpan(t *testing.T) {
	inputs := [][]int{
		[]int{1},
		[]int{1, 2, 1, 1, 3},
		[]int{1, 4, 2, 1, 4, 1, 4},
		[]int{1, 4, 2, 1, 4, 4, 4},
	}
	want := []int{1, 4, 6, 6}
	for index, input := range inputs {
		assert.Equal(t, want[index], maxSpan(input), "Max span between to array elements should be returned")
	}
}
