package algorithms

type UnionFind struct {
	parent []int
	rank []int
	size []int
	maxSize int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank: make([]int, n),
		size: make([]int, n),
		maxSize: 0,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int){
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY]{
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
			uf.maxSize = Max(uf.maxSize, uf.size[rootX])
		}
	}
}
