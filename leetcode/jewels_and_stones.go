package main

import "fmt"

//https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	j := make(map[rune]bool)
	for _, el := range J {
		j[el] = true
	}
	count := 0
	for _, el := range S {
		if j[el] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
