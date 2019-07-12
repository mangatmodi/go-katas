package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func TestPerimeter(t *testing.T) {
	r := &Rectangle{10.0, 10.}
	assert.Equal(t, r.Perimeter(), 40.0)
}
func TestArea(t *testing.T) {
	r := &Rectangle{10.0, 10.}
	assert.Equal(t, r.Area(), 100.0)
}

func TestCArea(t *testing.T) {
	c := &Circle{10.0}
	assert.Equal(t, c.Area(), 314.1592653589793)
}
