package unionfind

//MapUnionFind 并查集（map存储）
type MapUnionFind struct {
	parent map[int]int
	rank   map[int]int
	size   map[int]int
	count  int
}

//NewMapUnionFind 初始化带权重的并查集
func NewMapUnionFind(n int) *MapUnionFind {
	uf := &MapUnionFind{
		parent: make(map[int]int, n),
		rank:   make(map[int]int, n),
		size:   make(map[int]int, n),
	}
	return uf
}

//Union 合并两个节点,路径压缩，按秩合并
func (s *MapUnionFind) Union(x, y int) bool {
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
func (s *MapUnionFind) Find(x int) int {
	if _, ok := s.parent[x]; !ok {
		s.parent[x] = x
		s.count++
	}
	if x != s.parent[x] {
		s.parent[x] = s.Find(s.parent[x])
	}
	return s.parent[x]
}

//IsConnected 节点是否连通
func (s *MapUnionFind) IsConnected(x, y int) bool {
	return s.Find(x) == s.Find(y)
}

//Count ..
func (s *MapUnionFind) Count() int {
	return s.count
}

//Size ..
func (s *MapUnionFind) Size(x int) int {
	return s.size[x]
}
