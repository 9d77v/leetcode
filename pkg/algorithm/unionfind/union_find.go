package unionfind

//UnionFind 并查集
type UnionFind interface {
	Union(x, y int) bool
	Find(x int) int
	Size() int
	IsConnected(x, y int) bool
}
