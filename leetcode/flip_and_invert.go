package main

func flipAndInvertImage(A [][]int) [][]int {
	for i, arr := range A {
		A[i] = reverseAndFlip(arr)
	}
	return A
}

func reverseAndFlip(A []int) []int {
	i := 0
	j := len(A) - 1
	for i <= j {
		t := A[j]
		A[j] = flip(A[i])
		A[i] = flip(t)
		i++
		j--
	}
	return A
}

func flip(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}

func main() {
	//flipAndInvertImage()
}
