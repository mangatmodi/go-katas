package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	a := []int{3, 2, 1, 1, 4}
	assert.Equal(t, canJump(a), false)
}
