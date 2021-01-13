package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
在本问题中, 树指的是一个连通且无环的无向图。

输入一个图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N) 的树及一条附加的边构成。附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。

结果图是一个以边组成的二维数组。每一个边的元素是一对[u, v] ，满足 u < v，表示连接顶点u 和v的无向图的边。

返回一条可以删去的边，使得结果图是一个有着N个节点的树。如果有多个答案，则返回二维数组中最后出现的边。答案边 [u, v] 应满足相同的格式 u < v。
注意:

输入的二维数组大小在 3 到 1000。
二维数组中的整数在1到N之间，其中N是输入数组的大小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/redundant-connection
*/

/*
方法一：并查集
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：16 ms	内存消耗：6 MB
*/
func findRedundantConnectionFunc1(edges [][]int) []int {
	n := len(edges)
	unionFind := NewUnionFind(n + 1)
	for _, edge := range edges {
		if !unionFind.Union(edge[0], edge[1]) {
			return edge
		}
	}
	return []int{}
}

/*
方法二：拓扑排序
时间复杂度：О(nlogn)
空间复杂度：О(n)
运行时间：4 ms	内存消耗：3.5 MB
*/
func findRedundantConnectionFunc2(edges [][]int) []int {
	n := len(edges) + 1
	graph := make([][]int, n)
	indegrees := make([]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		indegrees[edge[0]]++
		indegrees[edge[1]]++
	}
	q := []int{}
	for i, indegree := range indegrees {
		if indegree == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range graph[u] {
			indegrees[v]--
			if indegrees[v] == 1 {
				q = append(q, v)
			}
		}
	}
	for i := n - 2; i >= 0; i-- {
		if indegrees[edges[i][0]] > 1 && indegrees[edges[i][1]] > 1 {
			return edges[i]
		}
	}
	return []int{}
}
