package main

import "fmt"

//https://leetcode.com/problems/prime-arrangements/submissions/
var primes = [25]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
var mod = 1000000007

func numPrimeArrangements(n int) int {
	nOfPrimes := 0
	for i := 0; i < 25 && primes[i] <= n; i++ {
		nOfPrimes++
	}
	return (factorial(nOfPrimes) * factorial(n-nOfPrimes)) % mod
}

func factorial(n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans = ((ans % mod) * (i % mod)) % mod
	}
	return ans
}

func main() {
	fmt.Println(numPrimeArrangements(100))
}
