package array

import "github.com/mangatmodi/go-katas/trees"

type Sorting interface {
	Sort(arr []int) []int
}

func Sort(arr []int, algo Sorting) []int {
	return algo.Sort(arr)
}

type MergeSort struct{}

func (m MergeSort) Sort(arr []int) []int {
	size := len(arr)
	if size < 2 {
		return arr
	}

	mid := len(arr) / 2

	arr1 := m.Sort(arr[:mid])
	arr2 := m.Sort(arr[mid:])

	return m.merge(arr1, arr2)
}

func (m MergeSort) merge(arr1 []int, arr2 []int) (merged []int) {
	size := len(arr1) + len(arr2)
	arr := make([]int, size)

	i := 0
	j := 0
	k := 0
	for ; k < size && j < len(arr2) && i < len(arr1); k++ {
		if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			i++
		} else {
			arr[k] = arr2[j]
			j++
		}
	}
	for ; k < size && j < len(arr2); k++ {
		arr[k] = arr2[j]
		j++
	}

	for ; k < size && i < len(arr1); k++ {
		arr[k] = arr1[i]
		i++
	}
	return arr
}

type QuickSort struct{}

func (q QuickSort) Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	p := q.Partition(arr, len(arr)/2)
	if p > 1 {
		q.Sort(arr[:p])
	}
	if p < len(arr)-1 {
		q.Sort(arr[p+1:])
	}

	return arr
}

func (q QuickSort) Partition(arr []int, idx int) int {
	if len(arr) < 2 {
		return idx
	}
	start := 0
	end := len(arr) - 1

	v := arr[idx]
	for end >= start {
		if arr[end] <= v {
			t := arr[end]
			arr[end] = v
			arr[idx] = t
			idx = end
		} else {
			end--
		}
		if arr[start] > v {
			t := arr[start]
			arr[start] = v
			arr[idx] = t
			idx = start
		} else {
			start++
		}
	}
	return idx
}

type TreeSort struct{}

func (t TreeSort) Sort(arr []int) []int {
	return trees.Infix(trees.ToBST(arr))
}
