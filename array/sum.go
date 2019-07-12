package array

func Sum(a []int) int32 {
	var s int32
	for _, n := range a {
		s += int32(n)
	}
	return s
}
