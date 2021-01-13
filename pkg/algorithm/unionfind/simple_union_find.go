package unionfind

//SimpleUnionFind 简单并查集
type SimpleUnionFind struct {
	parent []int
	rank   []int
	size   int
}

//NewSimpleUnionFind 初始化带权重的并查集
func NewSimpleUnionFind(n int) *SimpleUnionFind {
	uf := &SimpleUnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

//Union 合并两个节点,路径压缩，按秩合并
func (s *SimpleUnionFind) Union(x, y int) bool {
	rootX, rootY := s.Find(x), s.Find(y)
	if rootX != rootY {
		if s.rank[rootX] == s.rank[rootY] {
			s.parent[rootX] = rootY
			s.rank[rootY]++
		} else if s.rank[rootX] < s.rank[rootY] {
			s.parent[rootX] = rootY
		} else {
			s.parent[rootY] = rootX
		}
		s.size--
		return true
	}
	return false
}

//Find 路径压缩
func (s *SimpleUnionFind) Find(x int) int {
	if x != s.parent[x] {
		s.parent[x] = s.Find(s.parent[x])
	}
	return s.parent[x]
}

//IsConnected 节点是否连通
func (s *SimpleUnionFind) IsConnected(x, y int) bool {
	rootX, rootY := s.Find(x), s.Find(y)
	return rootX == rootY
}

//Size ..
func (s *SimpleUnionFind) Size() int {
	return s.size
}
