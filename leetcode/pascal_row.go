package main

import (
	"encoding/json"
	"fmt"
)

//https://leetcode.com/problems/pascals-triangle-ii/

func getRow(rowIndex int) []int {
	switch rowIndex {
	case 0:
		return []int{1}
	case 1:
		return []int{1, 1}
	}

	prev := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		a := make([]int, i+1)
		a[0] = 1
		a[len(a)-1] = 1
		for j := 1; j < len(a)-1; j++ {
			a[j] = prev[j-1] + prev[j]
		}
		prev = a
	}
	return prev
}

func main() {
	a := getRow(6)
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}
