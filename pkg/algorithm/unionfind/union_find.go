package unionfind

//RankType 按秩合并类型
type RankType int

//按秩合并类型
const (
	RankNone   RankType = iota //不按秩合并
	RankHeight                 //按树的高度合并
	RankSize                   //按树的节点数合并
)

//UnionFind 并查集
type UnionFind interface {
	Union(x, y int) bool
	Find(x int) int
	Rank(x int) int
	Count() int
	IsConnected(x, y int) bool
}
