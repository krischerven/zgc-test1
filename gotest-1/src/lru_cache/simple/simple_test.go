package simple

import (
	"fmt"
	"github.com/krischerven/zgc-test1/gotest-1/src/util/test"
	"testing"
)

func TestLRUcache(t *testing.T) {
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
	test.Assert(
		t,
		c.is(new(5), new(4), new(1), new(3)),
		c.elements(),
		[]int{5, 4, 1, 3},
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
	for _, i := range []int{5, 4, 1, 3} {
		test.Assert(
			t,
			c.Hit(&i) == true,
			fmt.Sprintf("%t (%d)", false, i),
			fmt.Sprintf("%t (%d)", true, i),
		)
	}
	for _, i := range []int{2, 6, 7, 8} {
		test.Assert(
			t,
			c.Hit(&i) == false,
			fmt.Sprintf("%t (%d)", true, i),
			fmt.Sprintf("%t (%d)", false, i),
		)
	}
}
