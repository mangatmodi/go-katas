package main

import (
	"fmt"
)

// Min number of palindrom strings in a string
func removePalindromeSub(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 0

	for i := range s {
		//i is prev in dp
		c := cost(s, dp, i)
		dp[i+1] = min((1 + dp[i]), c)
	}
	return dp[len(s)]
}

func cost(s string, dp []int, end int) int {
	l := end + 1
	for l > 1 {
		if isPalindrome(s, l, end) {
			return min(dp[end-l+1], dp[end-1+1]) + 1
		}
		l--
	}
	return dp[end-1+1] + 1
}

func isPalindrome(s string, l, end int) bool {
	i := end - l + 1
	j := end
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(removePalindromeSub("bbaabaaa"))
}
