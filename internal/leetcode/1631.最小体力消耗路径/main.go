package main

import (
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row, col) 的高度。一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。你每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。

一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。

请你返回从左上角走到右下角的最小 体力消耗值 。
提示：

rows == heights.length
columns == heights[i].length
1 <= rows, columns <= 100
1 <= heights[i][j] <= 10^6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-with-minimum-effort
*/

/*
方法一：并查集
时间复杂度：О(2*m*n-m-nlog(2*m*n-m-n))
空间复杂度：О(2*m*n-m-n)
运行时间：260 ms	内存消耗：7.8 MB
*/
func minimumEffortPathFuc1(heights [][]int) (result int) {
	m := len(heights)
	if m == 0 {
		return 0
	}
	n := len(heights[0])
	if n == 0 {
		return 0
	}
	edges := make([][3]int, 2*m*n-m-n)
	t := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+1 < m {
				edges[t] = [3]int{index(i, j, n), index(i+1, j, n), distance(heights, i, j, i+1, j)}
				t++
			}
			if j+1 < n {
				edges[t] = [3]int{index(i, j, n), index(i, j+1, n), distance(heights, i, j, i, j+1)}
				t++
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	var uf UnionFind = NewMapUnionFindWithRank(len(edges), RankSize)
	for _, edge := range edges {
		if uf.Union(edge[0], edge[1]) {
			result = Max(result, edge[2])
			if uf.IsConnected(0, m*n-1) {
				return result
			}
		}
	}
	return
}

func index(i, j, n int) int {
	return i*n + j
}

func distance(heights [][]int, i1, j1, i2, j2 int) int {
	a, b := heights[i1][j1], heights[i2][j2]
	if a > b {
		return a - b
	}
	return b - a
}
