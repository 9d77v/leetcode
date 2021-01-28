package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
给定一个包含了一些 0 和 1 的非空二维数组 grid 。

一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

注意: 给定的矩阵grid 的长度和宽度都不超过 50。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-area-of-island
*/

/*
方法一：并查集(数组)
时间复杂度：О(nmlognm)
空间复杂度：О(nm)
运行时间：8 ms	内存消耗：5.6 MB
*/
func maxAreaOfIsland(grid [][]int) (result int) {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var uf UnionFind = NewArrayUnionFindWithRank(m*n, RankSize)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				current, right, down := index(i, j, m), index(i, j+1, m), index(i+1, j, m)
				if j+1 < m && grid[i][j+1] == 1 {
					uf.Union(current, right)
				}
				if i+1 < n && grid[i+1][j] == 1 {
					uf.Union(current, down)
				}
				result = Max(result, uf.Rank(uf.Find(current)))
			}
		}
	}
	return
}

func index(i, j, m int) int {
	return i*m + j
}

/*
方法二：深度优先遍历
时间复杂度：О(nm)
空间复杂度：О(nm)
运行时间：12 ms	内存消耗：5 MB
*/
func maxAreaOfIslandFunc2(grid [][]int) (result int) {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i, j-1) + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j+1)
	}
	for i := range grid {
		for j := range grid[0] {
			result = Max(result, dfs(i, j))
		}
	}
	return
}

/*
方法三：广度优先遍历
时间复杂度：О(nm)
空间复杂度：О(nm)
运行时间：12 ms	内存消耗：6.8 MB
*/
func maxAreaOfIslandFunc3(grid [][]int) (result int) {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var bfs func(i, j int) int
	bfs = func(i, j int) (count int) {
		arr := [][2]int{{i, j}}
		for len(arr) != 0 {
			a, b := arr[0][0], arr[0][1]
			arr = arr[1:]
			if a >= 0 && a < n && b >= 0 && b < m && grid[a][b] == 1 {
				count++
				grid[a][b] = 0
				arr = append(arr, [][2]int{{a - 1, b}, {a, b - 1}, {a + 1, b}, {a, b + 1}}...)
			}
		}
		return
	}
	for i := range grid {
		for j := range grid[0] {
			result = Max(result, bfs(i, j))
		}
	}
	return
}
