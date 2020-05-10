package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	sb := strings.Builder{}
	if len(strs) == 0 {
		return ""
	}
	for i := 0; ; i++ {
		var prev rune = -1
		for _, s := range strs {
			if i >= len(s) {
				return sb.String()
			}
			if prev != -1 {
				if prev != rune(s[i]) {
					return sb.String()
				}
			}
			prev = rune(s[i])
		}
		sb.WriteRune(rune(prev))
	}
//	return sb.String()
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{""}))
	fmt.Println(longestCommonPrefix([]string{"ab", "ac"}))

}
