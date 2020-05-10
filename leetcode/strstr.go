package main

import "fmt"

//https://leetcode.com/problems/implement-strstr/

//Trying boyer-moore string algorithm
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	pat := map[rune]int{}
	for i, r := range needle {
		pat[r] = i
	}

	l := len(needle) - 1
	for l < len(haystack) {
		match := true
		for i := 0; i <= len(needle)-1; i++ {
			if needle[len(needle)-i-1] != haystack[l-i] {
				l = l - i
				match = false
				break
			}
		}
		if match {
			return l - len(needle) + 1
		} else {
			idx, ok := pat[rune(haystack[l])]
			if !ok {
				l = l + len(needle)
			} else {
				if (len(needle) - idx - 1) == 0 {
					l = l + 1
				} else {
					l = l + (len(needle) - idx - 1)
				}
			}
		}
	}
	return -1
}

func main() {
	//fmt.Println(strStr("aababcabcdabcdeabcdef", "abcdef"))
	//fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}
