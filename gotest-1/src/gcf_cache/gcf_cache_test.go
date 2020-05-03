package gcf_cache

import (
	"fmt"
	"github.com/krischerven/zgc-test1/gotest-1/src/util/test"
	"testing"
)

func TestGCFcache(t *testing.T) {
	c := New(4)
	new := func(i int) Key {
		return Key{i}
	}
	c.Refer(new(1))
	{ // ensure repeat keys are not allowed
		err := c.Refer(new(1))
		test.Assert(t, err != nil, nil, err)
	}
	c.Refer(new(2))
	c.Refer(new(3))
	c.Refer(new(4))
	c.Refer(new(5))
	c.Refer(new(6))
	c.Refer(new(7))
	c.Refer(new(8))
	c.Refer(new(9))
	test.Assert(
		t,
		c.is(new(6), new(7), new(8), new(9)),
		c.elements(),
		[]Key{new(6), new(7), new(8), new(9)},
	)
	test.Assert(
		t,
		c.Cap() == 4,
		c.Cap(),
		4,
	)
	test.Assert(
		t,
		c.Size() == 4,
		c.Size(),
		4,
	)
	for _, i := range []int{6, 7, 8, 9} {
		test.Assert(
			t,
			c.Hit(Key{i}) == true,
			fmt.Sprintf("%t (%d)", false, i),
			fmt.Sprintf("%t (%d)", true, i),
		)
	}
	for _, i := range []int{1, 2, 3, 4} {
		test.Assert(
			t,
			c.Hit(Key{i}) == false,
			fmt.Sprintf("%t (%d)", true, i),
			fmt.Sprintf("%t (%d)", false, i),
		)
	}
}
