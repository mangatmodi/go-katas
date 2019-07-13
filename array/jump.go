package array

//https://leetcode.com/problems/jump-game/
/*
Make an array a[n] possible to reach from origin
or DFS/BFS from origin - still n2 as it is a graph
[3,2,2,0,1]
*/

func canJump(nums []int) bool {
	var seen = make(map[int]bool)
	var result = make([]bool, len(nums))
	if len(nums) < 2 {
		return true
	}
	return jump(nums, 0, seen, result)
}

func jump(nums []int, i int, seen map[int]bool, result []bool) bool {
	if i == len(nums)-1 {
		return true
	}
	if nums[i] == 0 {
		return false
	}
	if seen[i] {
		return result[i]
	}
	r := false
	for j := 1; j < nums[i] && j < len(nums); j++ {
		r = r || jump(nums, i+j, seen, result)
	}
	seen[i] = true
	result[i] = r
	return r
}
