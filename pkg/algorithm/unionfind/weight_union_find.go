package unionfind

//WeightUnionFind 带权重的并查集
type WeightUnionFind struct {
	parent []int
	weight []float64
	size   int
}

//NewWeightUnionFind 初始化带权重的并查集
func NewWeightUnionFind(n int) *WeightUnionFind {
	wuf := &WeightUnionFind{
		parent: make([]int, n),
		weight: make([]float64, n),
		size:   n,
	}
	for i := 0; i < n; i++ {
		wuf.parent[i] = i
		wuf.weight[i] = 1
	}
	return wuf
}

//Union 合并两个节点
func (w *WeightUnionFind) Union(x, y int, value float64) {
	rootX, rootY := w.find(x), w.find(y)
	if rootX != rootY {
		w.parent[rootX] = rootY
		w.weight[rootX] = w.weight[y] * value / w.weight[x]
		w.size--
	}
}

//路径压缩
func (w *WeightUnionFind) find(x int) int {
	if x != w.parent[x] {
		origin := w.parent[x]
		w.parent[x] = w.find(w.parent[x])
		w.weight[x] *= w.weight[origin]
	}
	return w.parent[x]
}

//IsConnected 节点是否连通
func (w *WeightUnionFind) IsConnected(x, y int) float64 {
	rootX, rootY := w.find(x), w.find(y)
	if rootX == rootY {
		return w.weight[x] / w.weight[y]
	}
	return -1
}

//Size ..
func (w *WeightUnionFind) Size() int {
	return w.size
}
