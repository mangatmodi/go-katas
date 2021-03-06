package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	sorted := Sort([]int{5, 1, 218, 9}, TreeSort{})
	assert.Equal(t, []int{1, 5, 9, 218}, sorted)
}

//100 integers
func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(
			[]int{
				1, 8, 44, 49, 50, 51, 68, 82, 87, 90, 95, 99, 104, 123, 136, 148, 180,
				188, 218, 226, 229, 234, 260, 281, 282, 300, 304, 307, 326, 331, 343, 353,
				360, 369, 372, 405, 418, 422, 427, 428, 435, 442, 453, 466, 473, 493, 502,
				503, 557, 565, 578, 587, 590, 591, 595, 607, 609, 611, 612, 616, 618, 638,
				641, 646, 648, 684, 707, 710, 726, 729, 752, 767, 772, 779, 788, 803, 825,
				832, 833, 848, 855, 861, 862, 863, 865, 869, 892, 894, 901, 902, 908, 918,
				920, 929, 943, 948, 953, 975, 982, 998,
			},
			MergeSort{},
		)
	}
}
