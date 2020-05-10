package main

import "fmt"

//https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	uniq := make(map[rune]bool)
	dup := make(map[rune]bool)
	ans := -1
	for _, r := range s {
		if !dup[r] {
			if uniq[r] {
				//not uniq anymore
				delete(uniq, r)
				dup[r] = true
			} else {
				uniq[r] = true
			}
		}
	}

	for i, r := range s {
		if uniq[r] {
			return i
		}
	}
	return ans
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
