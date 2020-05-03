package gcf_cache

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, b bool, have interface{}, want interface{}) {
	if !b {
		t.Errorf("Assertion failed! have: %v, want %v", have, want)
	}
}

func TestGCFcache(t *testing.T) {
	c := New(4)
	new := func(i int) Key {
		return Key{i}
	}
	// repeat keys don't work
	c.Refer(new(1))
	c.Refer(new(2))
	c.Refer(new(3))
	c.Refer(new(4))
	c.Refer(new(5))
	c.Refer(new(6))
	assert(
		t,
		c.is(new(3), new(4), new(5), new(6)),
		c.elements(),
		[]Key{new(3), new(4), new(5), new(6)},
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
	for _, i := range []int{3, 4, 5, 6} {
		assert(
			t,
			c.Hit(Key{i}) == true,
			fmt.Sprintf("%t (%d)", false, i),
			fmt.Sprintf("%t (%d)", true, i),
		)
	}
	for _, i := range []int{1, 2, 8, 9} {
		assert(
			t,
			c.Hit(Key{i}) == false,
			fmt.Sprintf("%t (%d)", true, i),
			fmt.Sprintf("%t (%d)", false, i),
		)
	}
}
