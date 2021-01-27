package unionfind

//ArrayUnionFind 并查集（数组存储）
type ArrayUnionFind struct {
	parent   []int
	rank     []int
	count    int
	size     int
	cap      int
	rankType RankType
}

//NewArrayUnionFind 初始化并查集
func NewArrayUnionFind(n int) *ArrayUnionFind {
	uf := &ArrayUnionFind{
		parent: make([]int, n),
		cap:    n,
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
		cap:      n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = -1
		uf.rank[i] = 1
	}
	return uf
}

//Union 合并两个结点,按秩合并
func (uf *ArrayUnionFind) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		switch uf.rankType {
		case RankHeight:
			if uf.rank[rootX] == uf.rank[rootY] {
				uf.rank[rootY]++
			} else if uf.rank[rootX] > uf.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
		case RankSize:
			if uf.rank[rootX] >= uf.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
			uf.rank[rootY] += uf.rank[rootX]
		}
		uf.parent[rootX] = rootY
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
		uf.size++
	}
	if x != uf.parent[x] {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

//IsConnected 结点是否连通
func (uf *ArrayUnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

//Count 连通树的数量
func (uf *ArrayUnionFind) Count() int {
	return uf.count
}

//Size 结点总数
func (uf *ArrayUnionFind) Size() int {
	return uf.size
}

//Rank 结点所在树的高度或结点数
func (uf *ArrayUnionFind) Rank(x int) int {
	return uf.rank[x]
}

//Has 是否存在某一结点
func (uf *ArrayUnionFind) Has(x int) bool {
	return uf.parent[x] != -1
}

//Cap 初始化容量
func (uf *ArrayUnionFind) Cap() int {
	return uf.cap
}
