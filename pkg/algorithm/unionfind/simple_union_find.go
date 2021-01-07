package unionfind

//SimpleUnionFind 简单并查集
type SimpleUnionFind struct {
	parent []int
	size   int
}

//NewSimpleUnionFind 初始化带权重的并查集
func NewSimpleUnionFind(n int) *SimpleUnionFind {
	uf := &SimpleUnionFind{
		parent: make([]int, n),
		size:   n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

//Union 合并两个节点
func (s *SimpleUnionFind) Union(x, y int) {
	rootX, rootY := s.find(x), s.find(y)
	if rootX != rootY {
		s.parent[rootX] = rootY
		s.size--
	}
}

//路径压缩
func (s *SimpleUnionFind) find(x int) int {
	if x != s.parent[x] {
		s.parent[x] = s.find(s.parent[x])
	}
	return s.parent[x]
}

//IsConnected 节点是否连通
func (s *SimpleUnionFind) IsConnected(x, y int) float64 {
	rootX, rootY := s.find(x), s.find(y)
	if rootX == rootY {
		return 1
	}
	return -1
}

//Size ..
func (s *SimpleUnionFind) Size() int {
	return s.size
}
