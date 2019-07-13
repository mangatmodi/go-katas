package string

import (
	"strings"
)

//https://leetcode.com/problems/reverse-words-in-a-string-iii/
func reverseWords(s string) string {
	sb := strings.Builder{}
	temp := strings.Builder{}
	noSpace := true
	for _, c := range s {
		if c == ' ' {
			if noSpace {
				add(&sb, &temp)
				noSpace = false
			}
			sb.WriteRune(' ')
			continue
		}
		if !noSpace {
			noSpace = true
		}
		temp.WriteRune(c)
	}

	add(&sb, &temp)
	return sb.String()
}
func add(sb *strings.Builder, temp *strings.Builder) {
	str := temp.String()
	temp.Reset()
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}
}
