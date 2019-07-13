package string

import "strings"

//https://leetcode.com/problems/add-binary/

func addBinary(a string, b string) string {
	var short, big = smallToBig(a, b)
	var lenS, lenB = len(short), len(big)
	sb := strings.Builder{}

	c := '0'
	for i := 0; i < lenB; i++ {
		var r, b1, b2 rune
		b1 = rune(big[lenB-i-1])
		if lenS-i-1 >= 0 {
			b2 = rune(short[lenS-i-1])
		} else {
			b2 = '0'
		}
		r, c = addBits(b1, b2, c)
		sb.WriteRune(r)
	}
	if c == '1' {
		sb.WriteRune(c)
	}
	return reverse(sb.String())

}

func reverse(s string) string {
	sb := strings.Builder{}
	for i := len(s) - 1; i > -1; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func smallToBig(a string, b string) (short string, big string) {
	if len(a) < len(b) {
		return a, b
	}
	return b, a
}

func addBits(b1 rune, b2 rune, c rune) (b3 rune, carry rune) {
	b3 = '0'
	carry = '0'
	switch {
	case b1 == '1' && b2 == '1':
		b3 = c
		carry = '1'
	case b1 == '0' && b2 == '0':
		b3 = c
		carry = '0'
	default:
		b3 = '1'
		if c == '1' {
			b3 = '0'
			carry = '1'
		}

	}
	return b3, carry
}
