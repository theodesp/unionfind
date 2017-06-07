package unionfind

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

// Package unionfind implements a weighted Union Find data structure with path compression.

// Union-Find is used to determine the connected components in a graph.
// We can determine whether 2 nodes are in the same connected component or not in the graph.
// We can also determine that by adding an edge between 2 nodes whether it leads to cycle in the graph or not.

//// Any type that implements unionfind.Interface can be used as
//type Interface interface {
//	Less(i, j int) bool
//	GetItem(i int) interface{}
//	SetItem(i int)
//}

//
type UnionFind struct {
	root []int
	size []int
}

// New returns an initialized list of Size N
func New(N int) *UnionFind {
	return new(UnionFind).init(N)
}

// Constructor initializes root and size arrays
func (uf *UnionFind) init(N int) *UnionFind {
	uf = new(UnionFind)
	uf.root = make([]int, N)
	uf.size = make([]int, N)

	for i := 0; i < N; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}

	return uf
}

// Union connects p and q by finding their roots and comparing their respective
// size arrays to keep the tree flat
func (uf *UnionFind) Union(p int, q int) {
	qRoot := uf.Root(q)
	pRoot := uf.Root(p)

	if uf.size[qRoot] < uf.size[pRoot] {
		uf.root[qRoot] = uf.root[pRoot]
		uf.size[pRoot] += uf.size[qRoot]
	} else {
		uf.root[pRoot] = uf.root[qRoot]
		uf.size[qRoot] += uf.size[pRoot]
	}
}

// Root or Find traverses each parent element while compressing the
// levels to find the root element of p
// If we attempt to access an element outside the array it returns -1
func (uf *UnionFind) Root(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}

	for uf.root[p] != p {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}

	return p
}

// Root or Find
func (uf *UnionFind) Find(p int) int {
	return uf.Root(p)
}

// Check if items p,q are connected
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Root(p) == uf.Root(q)
}
