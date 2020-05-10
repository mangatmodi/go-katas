package main

import (
	"encoding/json"
	"fmt"
)

//https://leetcode.com/problems/permutations/
func permute(nums []int) [][]int {
	all := make([][]int, 0, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		//start from i
		perm := make([]int, len(nums))
		perm[0] = nums[i]
		rec(nums, perm, 1, i, (i+1)%len(nums))
		all = append(all, perm)
	}
	return all
}

func rec(nums []int, perm []int, idx, start, curr int) {
	if start == curr {
		return
	}
	if idx >= len(nums) {
		return
	}
	perm[idx] = nums[curr]
	rec(nums, perm, idx+1, start, (curr+1)%len(nums))
}

func main() {
	a := []int{1, 2, 3}
	ans := permute(a)
	b, _ := json.Marshal(ans)
	fmt.Println(string(b))
}
