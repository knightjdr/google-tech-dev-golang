// maxspan is my solution to https://codingbat.com/prob/p189576
package maxspan

// maxSpan works for slices of any type
func maxSpan(arr []interface{}) int {
	// Find first and last index for each value in the slice.
	span := make(map[interface{}][]int)
	for index, value := range arr {
		if _, ok := span[value]; ok {
			span[value][1] = index
		} else {
			span[value] = make([]int, 2)
			span[value][0] = index
		}
	}

	// Find the greatest span between values, including last occurrence.
	max := 1
	for _, tuple := range span {
		if len(tuple) == 2 {
			tupleSpan := tuple[1] - tuple[0] + 1
			if tupleSpan > max {
				max = tupleSpan
			}
		}
	}
	return max
}

// maxSpanInt only works for integer slices
func maxSpanInt(arr []int) int {
	// Find first and last index for each value in the slice.
	span := make(map[int][]int)
	for index, value := range arr {
		if _, ok := span[value]; ok {
			span[value][1] = index
		} else {
			span[value] = make([]int, 2)
			span[value][0] = index
		}
	}

	// Find the greatest span between values, including last occurrence.
	max := 1
	for _, tuple := range span {
		if len(tuple) == 2 {
			tupleSpan := tuple[1] - tuple[0] + 1
			if tupleSpan > max {
				max = tupleSpan
			}
		}
	}
	return max
}
