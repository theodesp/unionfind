package unionfind

import (
	"sync"
)

/*
Open Source Initiative OSI - The MIT License (MIT):Licensing
The MIT License (MIT)
Copyright (c) 2017 Theo Despoudis (thdespou@hotmail.com)
Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Thread safe version of Union-Find using plain locks.


type ThreadSafeUnionFind struct {
	sync.RWMutex
	uf *UnionFind
}

func NewThreadSafeUnionFind(size int) ThreadSafeUnionFind {
	safeUnionFind := ThreadSafeUnionFind{}
	safeUnionFind.uf = New(size)

	return safeUnionFind
}

func (suf *ThreadSafeUnionFind) Union(p int, q int) {
	suf.Lock()
	defer suf.Unlock()

	suf.uf.Union(p, q)
}


func (suf *ThreadSafeUnionFind) Root(p int) int {
	suf.Lock()
	defer suf.Unlock()

	return suf.uf.Root(p)
}

// Root or Find
func (suf *ThreadSafeUnionFind) Find(p int) int {
	return suf.Root(p)
}

// Unfortunately all the calls are coerced to writes thats why we use a Writer lock
func (suf *ThreadSafeUnionFind) Connected(p int, q int) bool {
	suf.Lock()
	defer suf.Lock()

	return suf.uf.Root(p) == suf.uf.Root(p)
}


