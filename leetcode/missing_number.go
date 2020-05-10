package main

import "fmt"

//https://leetcode.com/problems/missing-number/
func missingNumber(nums []int) int {
	n := len(nums)
	reqSum := (n * (n + 1)) / 2
	sum := 0
	for _, el := range nums {
		sum += el
	}
	return reqSum - sum
}

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
