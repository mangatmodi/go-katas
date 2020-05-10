package main

import (
	"fmt"
	"math"
)

//https://leetcode.com/problems/count-primes/

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	primes := make([]int, 1, n/2)
	primes[0] = 2 //first prime

	for i := 3; i < n; i++ {
		until := int(math.Ceil(math.Sqrt(float64(i))))
		isPrime := true
		for j := 0;j < len(primes) && primes[j] <= until ; j++ {
			if i%primes[j] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return len(primes)
}

func main() {
	fmt.Println(countPrimes(100))
}
