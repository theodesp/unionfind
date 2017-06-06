# unionfind
An idiomatic implementation of a weighted Union Find data structure with path compression in go.

## WIP
Work in progress

## Install

`$ go get github.com/theodesp/unionfind`

## Usage
```go
// Initialize with size
uf := unionfind.New(10)

// Union a,b connects components at index a and b
uf.Union(1, 2)
uf.Union(2, 3)
uf.Union(5, 6)
uf.Union(4, 6)

fmt.PrintLn(uf.Find(2)) // Prints 1
fmt.PrintLn(uf.Connected(1, 2)) // Prints true
fmt.PrintLn(uf.Connected(1, 3)) // Prints false

```

## API

#### `uf := unionfind.New(N)`
Create a new Union Find Structure with size N.

#### `uf.Union(p, q)`
Connects p and q components.

#### `uf.Find(p)`
Attempts to find the Root component of p. Returns -1 if p is out of bounds.

#### `uf.Root(p, q)`
Same as Find.

#### `uf.Connected(p, q)`
Checks if p and q are connected.


## Licence
MIT - Theo Despoudis