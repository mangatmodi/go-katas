package array

/**
* https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/
* [1,1,7]
* [1,2,9]
[0,2,1,-6,6,-7,9,1,2,0,1]
[0,2,3,-3,3,-4,5,6,8,8,9]

[0,2,1,-6,6,7,9,-1,2,0,1]
[0,2,3,-3,3,10,19,18,20,20,21]
*
*/

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, n := range A {
		sum += n
	}

	if sum%3 != 0 {
		return false
	}

	target := sum / 3
	mul := [3]int{1, 2} //We need to points where sum is target and 2 * target
	p := 0
	sum = 0
	for _, n := range A {
		sum += n
		if sum == mul[p]*target {
			p++
		}
		if p == 2 {
			return true
		}
	}

	return p == 2 //if never found the 2nd point
}
