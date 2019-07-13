package string

//https://leetcode.com/problems/find-the-difference/
func findTheDifference(s string, t string) byte {
	ans := byte(0)
	for _, c := range s {
		ans = ans ^ byte(c)
	}
	for _, c := range t {
		ans = ans ^ byte(c)
	}

	return ans //s is empty
}
