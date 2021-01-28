package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
*/

/*
方法一：并查集(数组)
时间复杂度：О(nmɑ(MN))
空间复杂度：О(nm)
运行时间：4 ms	内存消耗：3.6 MB
*/
func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var uf UnionFind = NewArrayUnionFindWithRank(m*n, RankSize)
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				cur, right, down := index(i, j, m), index(i, j+1, m), index(i+1, j, m)
				uf.Union(cur, cur)
				if j+1 < m && grid[i][j+1] == '1' {
					uf.Union(cur, right)
				}
				if i+1 < n && grid[i+1][j] == '1' {
					uf.Union(cur, down)
				}
			}
		}
	}
	return uf.Count()
}

func index(i, j, m int) int {
	return i*m + j
}
