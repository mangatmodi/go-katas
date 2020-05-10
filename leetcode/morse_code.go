package main

import (
	"fmt"
	"strings"
)

//https://leetcode.com/problems/unique-morse-code-words/

var morse = [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	set := map[string]struct{}{}
	for _, el := range words {
		sb := strings.Builder{}
		for _, r := range el {
			c := morse[r-97]
			sb.WriteString(c)
		}
		set[sb.String()] = struct{}{}
	}
	return len(set)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
