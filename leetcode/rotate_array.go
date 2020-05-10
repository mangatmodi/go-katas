package main

import (
	"encoding/json"
	"fmt"
)

//https://leetcode.com/problems/rotate-array/

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}

	changes := make(map[int]bool)
	i := 0
	el := nums[i]
	for len(changes) < len(nums) { //we need to change all
		if changes[i] { //Already put it at right place
			i = (i + 1) % len(nums)
			el = nums[i]
		}

		newI := (i + k) % len(nums)
		temp := nums[newI]
		nums[newI] = el
		changes[i] = true
		//Now put newI at the right value
		el = temp
		i = newI
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 17)
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}
