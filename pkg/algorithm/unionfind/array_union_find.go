package unionfind

//ArrayUnionFind 并查集（数组存储）
type ArrayUnionFind struct {
	parent   []int
	rank     []int
	count    int
	rankType RankType
}

//NewArrayUnionFind 初始化并查集
func NewArrayUnionFind(n int) *ArrayUnionFind {
	uf := &ArrayUnionFind{
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = -1
	}
	return uf
}

//NewArrayUnionFindWithRank 初始化带秩的并查集
func NewArrayUnionFindWithRank(n int, rankType RankType) *ArrayUnionFind {
	uf := &ArrayUnionFind{
		parent:   make([]int, n),
		rank:     make([]int, n),
		rankType: rankType,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = -1
		uf.rank[i] = 1
	}
	return uf
}

//Union 合并两个节点,按秩合并
func (uf *ArrayUnionFind) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		switch uf.rankType {
		case RankNone:
			uf.parent[rootX] = rootY
		case RankHeight:
			if uf.rank[rootX] == uf.rank[rootY] {
				uf.rank[rootY]++
			} else if uf.rank[rootX] > uf.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
			uf.parent[rootX] = rootY
		case RankSize:
			if uf.rank[rootX] >= uf.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
			uf.parent[rootX] = rootY
			uf.rank[rootY] += uf.rank[rootX]
		}
		uf.count--
		return true
	}
	return false
}

//Find 路径压缩
func (uf *ArrayUnionFind) Find(x int) int {
	if uf.parent[x] == -1 {
		uf.parent[x] = x
		uf.count++
	}
	if x != uf.parent[x] {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

//IsConnected 节点是否连通
func (uf *ArrayUnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

//Count ..
func (uf *ArrayUnionFind) Count() int {
	return uf.count
}

//Rank ..
func (uf *ArrayUnionFind) Rank(x int) int {
	return uf.rank[x]
}

//Has ..
func (uf *ArrayUnionFind) Has(x int) bool {
	return uf.parent[x] != -1
}
