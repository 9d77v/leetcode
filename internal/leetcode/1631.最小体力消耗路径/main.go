package main

import (
	"sort"

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
时间复杂度：О(mnlog(mn))
空间复杂度：О(mn)
运行时间：136 ms	内存消耗：7.2 MB
*/
func minimumEffortPathFuc1(heights [][]int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	m := len(heights[0])
	if m == 0 {
		return 0
	}
	edges := make([][3]int, 2*m*n-m-n)
	t := 0
	for i := range heights {
		for j := range heights[0] {
			if i > 0 {
				edges[t] = [3]int{index(i, j, m), index(i-1, j, m), distance(heights, i, j, i-1, j)}
				t++
			}
			if j > 0 {
				edges[t] = [3]int{index(i, j, m), index(i, j-1, m), distance(heights, i, j, i, j-1)}
				t++
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	var uf UnionFind = NewArrayUnionFindWithRank(m*n, RankSize)
	for _, edge := range edges {
		if uf.Union(edge[0], edge[1]) {
			if uf.IsConnected(0, m*n-1) {
				return edge[2]
			}
		}
	}
	return 0
}

func index(i, j, m int) int {
	return i*m + j
}

func distance(heights [][]int, i1, j1, i2, j2 int) int {
	a, b := heights[i1][j1], heights[i2][j2]
	if a > b {
		return a - b
	}
	return b - a
}

/*
方法二：二分查找+bfs
时间复杂度：О(mnlogC)
空间复杂度：О(mn)
运行时间：260 ms	内存消耗：7.8 MB
*/
func minimumEffortPathFuc2(heights [][]int) (result int) {
	n := len(heights)
	if n <= 0 {
		return 0
	}
	m := len(heights[0])
	if m <= 1 {
		return 0
	}
	var bfs func(target int) bool
	bfs = func(target int) bool {
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, m)
		}
		visited[0][0] = true
		q := [][2]int{{}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range directions {
				x, y := p[0]+d[0], p[1]+d[1]
				if x >= 0 && x < n && y >= 0 && y < m && !visited[x][y] &&
					distance(heights, x, y, p[0], p[1]) <= target {
					if x == n-1 && y == m-1 {
						return true
					}
					visited[x][y] = true
					q = append(q, [2]int{x, y})
				}
			}
		}
		return false
	}
	return sort.Search(1000000, func(mid int) bool {
		return bfs(mid)
	})
}

var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

/*
方法三：二分查找+dfs
时间复杂度：О(mnlogC)
空间复杂度：О(mn)
运行时间：260 ms	内存消耗：7.8 MB
*/
func minimumEffortPathFuc3(heights [][]int) (result int) {
	n := len(heights)
	if n <= 0 {
		return 0
	}
	m := len(heights[0])
	if m <= 1 {
		return 0
	}
	var dfs func(target, i, j int, visited [][]bool) bool
	dfs = func(target, i, j int, visited [][]bool) bool {
		visited[i][j] = true
		for _, d := range directions {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < n && y >= 0 && y < m && !visited[x][y] &&
				distance(heights, x, y, i, j) <= target {
				if x == n-1 && y == m-1 {
					return true
				}
				if dfs(target, x, y, visited) {
					return true
				}
			}
		}
		return false
	}
	return sort.Search(1000000, func(mid int) bool {
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, m)
		}
		visited[0][0] = true
		return dfs(mid, 0, 0, visited)
	})
}
