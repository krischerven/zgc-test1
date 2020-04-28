package simple

import (
	"testing"
)

func assert(t *testing.T, b bool, have interface{}, want interface{}) {
	if !b {
		t.Errorf("Assertion failed! have: %v, want %v", have, want)
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
	assert(
		t,
		c.is(new(5), new(4), new(1), new(3)),
		c.elements(),
		[]int{5, 4, 1, 3},
	)
	assert(
		t,
		c.Cap() == 4,
		c.Cap(),
		4,
	)
	assert(
		t,
		c.Size() == 4,
		c.Size(),
		4,
	)
}
