package array

func repeatedNTimes(A []int) int {
	m := make(map[int]bool, len(A)/2)
	for _, n := range A {
		if m[n] {
			return n
		}
		m[n] = true
	}
	return -1
}
