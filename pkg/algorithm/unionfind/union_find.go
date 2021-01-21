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
	Union(x, y int) bool       //按秩合并
	Find(x int) int            //路径压缩
	IsConnected(x, y int) bool //结点是否连通
	Count() int                //连通树的数量
	Size() int                 //结点总数
	Cap() int                  //初始化容量
	Rank(x int) int            //结点所在树的高度或结点数
	Has(x int) bool            //是否存在某一结点
}
