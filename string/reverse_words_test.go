package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "bba c   ba", reverseWords("abb c   ab"))
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWords("a  b c")
	}
}
