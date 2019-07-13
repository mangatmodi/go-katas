package array

func sortByEven(a []int) []int {
	if len(a) < 2 {
		return a
	}

	start := 0
	end := len(a) - 1

	even := -1
	odd := -1

	for start < end {
		if a[start]%2 == 0 {
			start++
		} else {
			odd = start
		}
		if a[end]%2 == 1 {
			end--
		} else {
			even = end
		}

		//Swap
		if odd != -1 && even != -1 {
			t := a[even]
			a[even] = a[odd]
			a[odd] = t
			odd = -1
			even = -1
		}
	}
	return a
}

// 111111111111
