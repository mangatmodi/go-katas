package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, Sum([]int{1, 4, 17, 6}), int32(28))
	assert.Equal(t, Sum([]int{}), int32(0))
}
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{6, 48, 473, 678922})
	}
}
