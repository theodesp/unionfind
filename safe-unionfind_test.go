package unionfind

import (
	. "gopkg.in/check.v1"
	"runtime"
	"math/rand"
	"sync"
)


var _ = Suite(&MySuite{})

const N = 100

func (s *MySuite) TestInitConcurrent(c *C) {
	runtime.GOMAXPROCS(2)
	ints := rand.Perm(N)

	var tsuf ThreadSafeUnionFind

	var wg sync.WaitGroup
	wg.Add(len(ints))

	for i := 0; i < len(ints); i++ {
		go func(i int) {
			tsuf = NewThreadSafeUnionFind(N)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for _, i := range ints {
		c.Assert(len(tsuf.uf.root), Equals, len(tsuf.uf.size), Commentf("#i == %d", i))
	}
}
