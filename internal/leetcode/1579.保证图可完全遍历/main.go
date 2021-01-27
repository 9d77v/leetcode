package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
Alice 和 Bob 共有一个无向图，其中包含 n 个节点和 3  种类型的边：

类型 1：只能由 Alice 遍历。
类型 2：只能由 Bob 遍历。
类型 3：Alice 和 Bob 都可以遍历。
给你一个数组 edges ，其中 edges[i] = [typei, ui, vi] 表示节点 ui 和 vi 之间存在类型为 typei 的双向边。请你在保证图仍能够被 Alice和 Bob 完全遍历的前提下，找出可以删除的最大边数。如果从任何节点开始，Alice 和 Bob 都可以到达所有其他节点，则认为图是可以完全遍历的。

返回可以删除的最大边数，如果 Alice 和 Bob 无法完全遍历图，则返回 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable
*/

/*
方法一：并查集
时间复杂度：О(m⋅α(n))
空间复杂度：О(m)
运行时间：328 ms	内存消耗 20.6 MB
*/
func maxNumEdgesToRemove(n int, edges [][]int) int {
	result := len(edges)
	var uf1 UnionFind = NewArrayUnionFindWithRank(n+1, RankSize)
	var uf2 UnionFind = NewArrayUnionFindWithRank(n+1, RankSize)
	for _, edge := range edges {
		if edge[0] == 3 && !uf1.IsConnected(edge[1], edge[2]) {
			uf1.Union(edge[1], edge[2])
			uf2.Union(edge[1], edge[2])
			result--
		}
	}
	for _, edge := range edges {
		switch edge[0] {
		case 1:
			if uf1.Union(edge[1], edge[2]) {
				result--
			}
		case 2:
			if uf2.Union(edge[1], edge[2]) {
				result--
			}
		}
	}
	if uf1.Count() == 1 && uf2.Count() == 1 && uf1.Size() == n && uf2.Size() == n {
		return result
	}
	return -1
}
