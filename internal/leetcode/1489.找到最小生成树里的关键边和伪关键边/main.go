package main

import (
	"math"
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
给你一个 n 个点的带权无向连通图，节点编号为 0 到 n-1 ，同时还有一个数组 edges ，其中 edges[i] = [fromi, toi, weighti] 表示在 fromi 和 toi 节点之间有一条带权无向边。最小生成树 (MST) 是给定图中边的一个子集，它连接了所有节点且没有环，而且这些边的权值和最小。

请你找到给定图中最小生成树的所有关键边和伪关键边。如果从图中删去某条边，会导致最小生成树的权值和增加，那么我们就说它是一条关键边。伪关键边则是可能会出现在某些最小生成树中但不会出现在所有最小生成树中的边。

请注意，你可以分别以任意顺序返回关键边的下标和伪关键边的下标。

提示：
2 <= n <= 100
1 <= edges.length <= min(200, n * (n - 1) / 2)
edges[i].length == 3
0 <= fromi < toi < n
1 <= weighti <= 1000
所有 (fromi, toi) 数对都是互不相同的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree
*/

/*
方法一：kruskal算法
时间复杂度：О(m2⋅α(n))
空间复杂度：О(m+n)
运行时间：48 ms	内存消耗 7 MB
*/
func findCriticalAndPseudoCriticalEdgesFunc1(n int, edges [][]int) [][]int {
	for i := range edges {
		edges[i] = append(edges[i], i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	mstValue := calcMST(NewArrayUnionFindWithRank(n, RankSize), -1, edges)
	result := make([][]int, 2)
	result[0] = []int{}
	result[1] = []int{}
	for i, edge := range edges {
		if calcMST(NewArrayUnionFindWithRank(n, RankSize), i, edges) > mstValue {
			result[0] = append(result[0], edge[3])
			continue
		}
		var uf2 UnionFind = NewArrayUnionFindWithRank(n, RankSize)
		uf2.Union(edge[0], edge[1])
		if edge[2]+calcMST(uf2, i, edges) == mstValue {
			result[1] = append(result[1], edge[3])
		}
	}
	return result
}

func calcMST(uf UnionFind, ignoredID int, edges [][]int) (mstValue int) {
	left := uf.Cap() - 1
	for i, edge := range edges {
		if i != ignoredID && uf.Union(edge[0], edge[1]) {
			mstValue += edge[2]
			left--
			if left == 0 {
				break
			}
		}
	}
	if uf.Count() > 1 || uf.Size() != uf.Cap() {
		return math.MaxInt64
	}
	return
}
