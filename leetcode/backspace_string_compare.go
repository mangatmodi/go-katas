package main

import "fmt"

//https://leetcode.com/problems/backspace-string-compare/

func backspaceCompare(S string, T string) bool {
	i := len(S) - 1
	j := len(T) - 1
	for i >= 0 && j >= 0 {
		if S[i] != '#' && T[j] != '#' {
			if S[i] == T[j] {
				i--
				j--
				continue
			} else {
				return false
			}
		}
		if S[i] == '#' {
			i = goNext(S, i)
		}
		if T[j] == '#' {
			j = goNext(T, j)
		}
	}
	if i >= 0 {
		i = goNext(S, i)
	}
	if j >= 0 {
		j = goNext(T, j)
	}

	return i < 0 && j < 0 //both are empty now
}

func goNext(S string, index int) int {
	//count #
	count := 0
	i := index
	for ; i >= 0; i-- {
		if S[i] == '#' {
			count++
		} else {
			count--
		}
		if count < 0 {
			return i
		}
	}
	return i - count // all was #
}

func main() {
	fmt.Print(backspaceCompare("#abc", "abc"))

	//fmt.Print(backspaceCompare("bxj##tw", "bxj###tw"))
}
