package main

import (
	"encoding/json"
	"fmt"
)

//https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	firstZero := first(nums, 0, func(i int) bool {
		return i == 0
	})
	if firstZero == -1 {
		//no zeroes
		return
	}

	firstNonZero := first(nums, firstZero, func(i int) bool {
		return i != 0
	})
	if firstNonZero == -1 {
		//all 0s
		return
	}

	//swap
	t := nums[firstZero]
	nums[firstZero] = nums[firstNonZero]
	nums[firstNonZero] = t

	moveZeroes(nums[firstZero+1:])
}

func first(nums []int, start int, cond func(int) bool) int {
	for i := start; i < len(nums); i++ {
		if cond(nums[i]) {
			return i
		}
	}
	return -1
}

func main() {
	a := []int{0, 2, 3, 4, 0, 6, 7}
	moveZeroes(a)
	b, _ := json.Marshal(a)
	fmt.Println(string(b))

}
