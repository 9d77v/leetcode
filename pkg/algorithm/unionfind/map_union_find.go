package unionfind

//MapUnionFind 并查集（map存储）
type MapUnionFind struct {
	parent   map[int]int
	rank     map[int]int
	count    int
	rankType RankType
}

//NewMapUnionFind 初始化并查集
func NewMapUnionFind() *MapUnionFind {
	uf := &MapUnionFind{
		parent: make(map[int]int, 0),
	}
	return uf
}

//NewMapUnionFindWithRank 初始化带秩的并查集
func NewMapUnionFindWithRank(rankType RankType) *MapUnionFind {
	uf := &MapUnionFind{
		parent:   make(map[int]int, 0),
		rank:     make(map[int]int, 0),
		rankType: rankType,
	}
	return uf
}

//Union 合并两个节点,按秩合并
func (uf *MapUnionFind) Union(x, y int) bool {
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
func (uf *MapUnionFind) Find(x int) int {
	if _, ok := uf.parent[x]; !ok {
		uf.parent[x] = x
		uf.count++
		if uf.rankType != RankNone {
			uf.rank[x] = 1
		}
	}
	if x != uf.parent[x] {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

//IsConnected 节点是否连通
func (uf *MapUnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

//Count ..
func (uf *MapUnionFind) Count() int {
	return uf.count
}

//Rank ..
func (uf *MapUnionFind) Rank(x int) int {
	return uf.rank[x]
}

//Has ..
func (uf *MapUnionFind) Has(x int) bool {
	_, ok := uf.parent[x]
	return ok
}
