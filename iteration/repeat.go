package iteration

import "strings"

//Repeat c n times
func Repeat(c string, n int) string {
	var str strings.Builder
	for i := 0; i < n; i++ {
		str.WriteString(c)
	}
	return str.String()
}
