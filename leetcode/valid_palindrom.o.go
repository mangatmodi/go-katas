package main

import "fmt"

//https://leetcode.com/problems/valid-palindrome-ii/

func validPalindrome(s string) bool {
	return isPalindrome(s, 0, len(s)-1, 0)
}

func isPalindrome(s string, i, j, count int) bool {
	if i >= j {
		return true
	}
	if s[i] == s[j] {
		return isPalindrome(s, i+1, j-1, count)
	}
	count++
	if count == 2 {
		return false
	}

	return isPalindrome(s, i+1, j, count) || isPalindrome(s, i, j-1, count)
}

func main() {
	fmt.Println(validPalindrome("abc"))
}
