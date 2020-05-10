package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lilysHomework(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want int32
	}{
		//{"all sorted", []int32{1, 2, 3, 4, 5, 6}, 0},
		//{"example case", []int32{7, 15, 12, 3}, 2},
		{"desc case", []int32{3, 4, 2, 5, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lilysHomework(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
