package main

import (
	"fmt"
	"math"
)

//https://leetcode.com/problems/koko-eating-bananas/

func minEatingSpeed(piles []int, H int) int {
	sum := total(piles)
	min := int(math.Ceil(float64(sum) / float64(H)))
	max := max(piles)
	if H == len(piles) {
		return int(max)
	}
	max = math.Min(math.Ceil(float64(sum)/float64(H-len(piles))), max)

	//now binary search on max and min
	return ans(piles, int32(min), int32(max), int32(H))
}

func ans(piles []int, min, max, h int32) int {
	mid := (max + min) / 2
	for max > min+1 {
		if valid(piles, mid, h) {
			max = mid //we can go lower
		} else {
			min = mid //we need to go higher
		}
		mid = (max + min) / 2
	}
	if valid(piles, min, h) {
		return int(min)
	} else {
		return int(max)
	}
}

func valid(piles []int, k, h int32) bool {
	var hrs int32 = 0
	for _, el := range piles {
		taken := math.Ceil(float64(el) / float64(k))
		hrs = hrs + int32(taken)
	}
	return hrs <= h
}

func max(a []int) float64 {
	m := float64(-1 << 31)
	for _, el := range a {
		m = math.Max(float64(el), m)
	}
	return m
}
func total(a []int) int64 {
	var s int64 = 0
	for _, el := range a {
		s += int64(el)
	}
	return s
}

func main() {
	fmt.Print(minEatingSpeed([]int{312884470}, 968709470))
}
