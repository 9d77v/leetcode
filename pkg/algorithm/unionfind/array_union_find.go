package unionfind

//ArrayUnionFind 并查集（数组存储）
type ArrayUnionFind struct {
	parent []int
	rank   []int
	size   []int
	count  int
}

//NewArrayUnionFind 初始化带权重的并查集
func NewArrayUnionFind(n int) *ArrayUnionFind {
	uf := &ArrayUnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = -1
		uf.rank[i] = 1
		uf.size[i] = 1
	}
	return uf
}

//Union 合并两个节点,路径压缩，按秩合并
func (s *ArrayUnionFind) Union(x, y int) bool {
	rootX, rootY := s.Find(x), s.Find(y)
	if rootX != rootY {
		if s.rank[rootX] == s.rank[rootY] {
			s.parent[rootX] = rootY
			s.rank[rootY]++
			s.size[rootY] += s.size[rootX]
		} else if s.rank[rootX] < s.rank[rootY] {
			s.parent[rootX] = rootY
			s.size[rootY] += s.size[rootX]
		} else {
			s.parent[rootY] = rootX
			s.size[rootX] += s.size[rootY]
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
	rootX, rootY := s.Find(x), s.Find(y)
	return rootX == rootY
}

//Count ..
func (s *ArrayUnionFind) Count() int {
	return s.count
}

//Size ..
func (s *ArrayUnionFind) Size(x int) int {
	return s.size[x]
}
