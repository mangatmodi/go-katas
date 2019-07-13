package array

import "sort"

/**
https://leetcode.com/problems/height-checker/
[1,1,4,2,1,3]
[1,1,1,2,3,4]
*/
func heightChecker(heights []int) int {
	var c int
	if len(heights) < 2 {
		return c
	}

	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)

	for i, n := range heights {
		if n != sorted[i] {
			c++
		}
	}

	return c
}
