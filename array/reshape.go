package array

/*
*  https://leetcode.com/problems/reshape-the-matrix/
 */

func matrixReshape(nums [][]int, r int, c int) [][]int {
	var count = len(nums) * len(nums[0])
	if (r * c) != count {
		return nums
	}
	var ret = make([][]int, r)

	var row = make([]int, c)
	var j = 0
	var i = 0
	for _, a := range nums {
		for _, el := range a {
			if j == c { //index
				ret[i] = row
				j = 0
				row = make([]int, c)
				i++
			}
			row[j] = el
			j++
		}
	}
	ret[i] = row
	return ret
}
