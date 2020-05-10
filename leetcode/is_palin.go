package main

import "fmt"

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; {
		a, b := s[i], s[j]

		switch {
		case s[i]-'A' < 26 && s[i]-'A' >= 0:
			a = 'a' + s[i] - 'A'
		case s[i]-'a' < 26 && s[i]-'a' >= 0:
			a = s[i]
		case s[i]-'0' < 10 && s[i]-'0' >= 0:
			a = s[i]
		default:
			i++
			continue
		}

		switch {
		case s[j]-'A' < 26 && s[j]-'A' >= 0:
			b = 'a' + s[j] - 'A'
		case s[j]-'a' < 26 && s[j]-'a' >= 0:
			b = s[j]
		case s[j]-'0' < 10 && s[j]-'0' >= 0:
			b = s[j]
		default:
			j--
			continue
		}

		if a != b {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("Ab,B a"))
}
