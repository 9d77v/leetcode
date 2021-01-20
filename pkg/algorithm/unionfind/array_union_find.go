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
func (s *ArrayUnionFind) Union(x, y int) bool {
	rootX, rootY := s.Find(x), s.Find(y)
	if rootX != rootY {
		switch s.rankType {
		case RankNone:
			s.parent[rootX] = rootY
		case RankHeight:
			if s.rank[rootX] == s.rank[rootY] {
				s.rank[rootY]++
			} else if s.rank[rootX] > s.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
			s.parent[rootX] = rootY
		case RankSize:
			if s.rank[rootX] >= s.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
			s.parent[rootX] = rootY
			s.rank[rootY] += s.rank[rootX]
		}
		s.count--
		return true
	}
	return false
}

//Find 路径压缩
func (s *ArrayUnionFind) Find(x int) int {
	if s.parent[x] == -1 {
		s.parent[x] = x
		s.count++
	}
	if x != s.parent[x] {
		s.parent[x] = s.Find(s.parent[x])
	}
	return s.parent[x]
}

//IsConnected 节点是否连通
func (s *ArrayUnionFind) IsConnected(x, y int) bool {
	return s.Find(x) == s.Find(y)
}

//Count ..
func (s *ArrayUnionFind) Count() int {
	return s.count
}

//Rank ..
func (s *ArrayUnionFind) Rank(x int) int {
	return s.rank[x]
}
