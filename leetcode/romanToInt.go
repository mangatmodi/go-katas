package main

import "fmt"

//https://leetcode.com/problems/roman-to-integer/
/**

  I can be placed before V (5) and X (10) to make 4 and 9.
  X can be placed before L (50) and C (100) to make 40 and 90.
  C can be placed before D (500) and M (1000) to make 400 and 900.

*/
func romanToInt(s string) int {
	ans := 0
	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case 'M':
			ans += 1000
			i++
		case 'V':
			ans += 5
			i++
		case 'X':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'C':
					ans += 90
					i += 2
				case 'L':
					ans += 40
					i += 2
				default:
					ans += 10
					i++
				}
			} else {
				ans += 10
				i++
			}
		case 'L':
			ans += 50
			i++
		case 'C':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'M':
					ans += 900
					i += 2
				case 'D':
					ans += 400
					i += 2
				default:
					ans += 100
					i++
				}
			} else {
				ans += 100
				i++
			}
		case 'D':
			ans += 500
			i++
		case 'I':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'X':
					ans += 9
					i += 2
				case 'V':
					ans += 4
					i += 2
				default:
					ans += 1
					i++
				}
			} else {
				ans += 1
				i++
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("IV"))
}
