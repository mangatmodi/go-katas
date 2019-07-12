package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	assert.Equal(t, []int{1}, Infix(Node{1, nil, nil}))
}

func TestBST(t *testing.T) {
	n := []int{4, 3, 2}
	b := ToBST(n)
	assert.Equal(t, []int{2, 3, 4}, Infix(b))
}

//Also used as sorting
func BenchmarkTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := []int{
			1, 8, 44, 49, 50, 51, 68, 82, 87, 90, 95, 99, 104, 123, 136, 148, 180,
			188, 218, 226, 229, 234, 260, 281, 282, 300, 304, 307, 326, 331, 343, 353,
			360, 369, 372, 405, 418, 422, 427, 428, 435, 442, 453, 466, 473, 493, 502,
			503, 557, 565, 578, 587, 590, 591, 595, 607, 609, 611, 612, 616, 618, 638,
			641, 646, 648, 684, 707, 710, 726, 729, 752, 767, 772, 779, 788, 803, 825,
			832, 833, 848, 855, 861, 862, 863, 865, 869, 892, 894, 901, 902, 908, 918,
			920, 929, 943, 948, 953, 975, 982, 998,
		}
		Infix(ToBST(n))
	}
}
