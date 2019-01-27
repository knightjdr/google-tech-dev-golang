// maxspan is my solution to https://codingbat.com/prob/p189576
package maxspan

func maxSpan(arr []int) int {
	// Find first and last index for each value in the array.
	span := make(map[int][]int)
	for index, value := range arr {
		if _, ok := span[value]; ok {
			span[value][1] = index
		} else {
			span[value] = make([]int, 2)
			span[value][0] = index
		}
	}

	// Find the greatest span between values.
	max := 1
	for _, tuple := range span {
		if len(tuple) == 2 {
			tupleSpan := tuple[1] + 1 - tuple[0]
			if tupleSpan > max {
				max = tupleSpan
			}
		}
	}
	return max
}
