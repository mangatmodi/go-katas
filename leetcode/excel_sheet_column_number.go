package main

import (
	"fmt"
	"math"
)

//https: //leetcode.com/problems/excel-sheet-column-number/

func titleToNumber(s string) int {
	ans := 0
	pow := len(s) - 1
	for _, r := range s {
		v := int(r - rune('A') + 1) // we need a to be 1
		mul := int(math.Pow(26, float64(pow)))
		ans += v * mul
		pow--
	}
	return ans
}

func main() {
	fmt.Println(titleToNumber("Z"))
}
