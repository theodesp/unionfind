package unionfind

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

// Test New
func (s *MySuite) TestNew(c *C) {
	uf := New(10)

	c.Assert(uf.size, DeepEquals, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	c.Assert(uf.root, DeepEquals, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

// Test Root when Empty
func (s *MySuite) TestRootEmpty(c *C) {
	uf := New(10)

	c.Assert(uf.Root(10), Equals, -1)
	c.Assert(uf.Root(9), Equals, 9)
}

// Test Root when It has connected components
func (s *MySuite) TestRootConnected(c *C) {
	uf := New(10)

	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(5, 6)
	uf.Union(2, 8)

	c.Assert(uf.Root(8), Equals, 2)
	c.Assert(uf.Root(2), Equals, 2)
	c.Assert(uf.Root(5), Equals, 6)
}

// Test Union when when we connect the same item
func (s *MySuite) TestUnionSame(c *C) {
	uf := New(10)

	uf.Union(2, 2)

	c.Assert(uf.Root(2), Equals, 2)
}

// Test Union checks size array when it assigns the root
func (s *MySuite) TestUnionChecksSizes(c *C) {
	uf := New(5)

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)

	c.Assert(uf.size, DeepEquals, []int{1, 4, 1, 1, 1})
}

// Test Components
func (s *MySuite) TestComponents(c *C) {
	uf := New(10)

	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(5, 6)
	uf.Union(2, 8)

	c.Assert(uf.Components(), Equals, 6)
}
