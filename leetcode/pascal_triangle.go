package main

import (
	"encoding/json"
	"fmt"
)

//https://leetcode.com/problems/pascals-triangle/

func generate(	numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		a := make([]int, i+1)
		a[0] = 1
		a[len(a)-1] = 1
		ans[i] = a
	}
	for i := 2; i < numRows; i++{
		a := ans[i]
		for j:=1; j<len(a)-1;j++{
			a[j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}

func main() {
	a := generate(0)
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}
