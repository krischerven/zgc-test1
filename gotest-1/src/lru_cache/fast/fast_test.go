package fast

import (
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Errorf("Assertion failed!")
	}
}

func TestLRUcache(t *testing.T) {
	// 5 4 1 3
	c := New(4)
	new := func(i int) *int {
		return &i
	}
	c.Refer(new(1))
	c.Refer(new(2))
	c.Refer(new(3))
	c.Refer(new(1))
	c.Refer(new(4))
	c.Refer(new(5))
	assert(t, c.Is(new(5), new(4), new(1), new(3)))
}