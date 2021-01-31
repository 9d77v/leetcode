package unionfind

//MapUnionFind 并查集（map存储）
type MapUnionFind struct {
	parent   map[int]int
	rank     map[int]int
	count    int
	cap      int
	rankType RankType
}

//NewMapUnionFind 初始化并查集
func NewMapUnionFind(n int, rankType ...RankType) *MapUnionFind {
	rt := RankSize
	if len(rankType) > 0 {
		rt = rankType[0]
	}
	uf := &MapUnionFind{
		parent:   make(map[int]int, n),
		cap:      n,
		rankType: rt,
	}
	if rt != RankNone {
		uf.rank = make(map[int]int, 0)
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

//IsConnected 结点是否连通
func (uf *MapUnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

//Count 连通树的数量
func (uf *MapUnionFind) Count() int {
	return uf.count
}

//Size 结点总数
func (uf *MapUnionFind) Size() int {
	return len(uf.parent)
}

//Rank 结点所在树的高度或结点数
func (uf *MapUnionFind) Rank(x int) int {
	return uf.rank[x]
}

//Has 是否存在某一结点
func (uf *MapUnionFind) Has(x int) bool {
	_, ok := uf.parent[x]
	return ok
}

//Cap 初始化容量
func (uf *MapUnionFind) Cap() int {
	return uf.cap
}
