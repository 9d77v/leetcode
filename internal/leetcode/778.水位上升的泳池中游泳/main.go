package main

import (
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
在一个 N x N 的坐标方格 grid 中，每一个方格的值 grid[i][j] 表示在位置 (i,j) 的平台高度。

现在开始下雨了。当时间为 t 时，此时雨水导致水池中任意位置的水位为 t 。你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两个平台。假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。

你从坐标方格的左上平台 (0，0) 出发。最少耗时多久你才能到达坐标方格的右下平台 (N-1, N-1)？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swim-in-rising-water
*/

/*
方法一：二分+bfs
时间复杂度：О(n2logn)
空间复杂度：О(n2)
运行时间：8 ms	内存消耗：6.8MB
*/
func swimInWaterFunc1(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var bfs func(mid int, visited [][]bool) bool
	bfs = func(mid int, visited [][]bool) bool {
		q := [][2]int{{0, 0}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range directions {
				x, y := p[0]+d[0], p[1]+d[1]
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] <= mid && !visited[x][y] {
					q = append(q, [2]int{x, y})
					if x == m-1 && y == n-1 {
						return true
					}
					visited[x][y] = true
				}
			}
		}
		return false
	}
	return sort.Search(n*n, func(mid int) bool {
		if grid[0][0] > mid {
			return false
		}
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, m)
		}
		visited[0][0] = true
		return bfs(mid, visited)
	})
}

func index(i, j, m int) int {
	return i*m + j
}

var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

/*
方法二：二分+dfs
时间复杂度：О(n2logn)
空间复杂度：О(n2)
运行时间：8 ms	内存消耗：4.7MB
*/
func swimInWaterFunc2(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var dfs func(mid, i, j int, visited [][]bool) bool
	dfs = func(mid, i, j int, visited [][]bool) bool {
		visited[i][j] = true
		for _, d := range directions {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] <= mid && !visited[x][y] {
				if x == n-1 && y == m-1 {
					return true
				}
				if dfs(mid, x, y, visited) {
					return true
				}
			}
		}
		return false
	}
	return sort.Search(n*n, func(mid int) bool {
		if grid[0][0] > mid {
			return false
		}
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, m)
		}
		return dfs(mid, 0, 0, visited)
	})
}

/*
方法三：二分+并查集
时间复杂度：О(n2logn)
空间复杂度：О(n2)
运行时间：16 ms	内存消耗：6.6MB
*/
func swimInWaterFunc3(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var bfs func(mid int) bool
	bfs = func(mid int) bool {
		var uf UnionFind = NewArrayUnionFindWithRank(n*n, RankSize)
		for i := range grid {
			for j := range grid[0] {
				if grid[i][j] <= mid {
					uf.Union(index(i, j, m), index(i, j, m))
					if i > 0 && grid[i-1][j] <= mid {
						uf.Union(index(i, j, m), index(i-1, j, m))
					}
					if j > 0 && grid[i][j-1] <= mid {
						uf.Union(index(i, j, m), index(i, j-1, m))
					}
				}
			}
		}
		return uf.IsConnected(0, index(n-1, m-1, m))
	}
	return sort.Search(n*n, func(mid int) bool {
		if grid[0][0] > mid {
			return false
		}
		return bfs(mid)
	})
}

/*
方法四：并查集
时间复杂度：О(n2logn)
空间复杂度：О(n2)
运行时间：12 ms	内存消耗：4.2MB
*/
func swimInWaterFunc4(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	data := make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			data[grid[i][j]] = index(i, j, m)
		}
	}
	var uf UnionFind = NewArrayUnionFindWithRank(n*n, RankSize)
	for num, value := range data {
		i, j := value/n, value%n
		for _, d := range directions {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] <= num {
				uf.Union(value, index(x, y, m))
				if uf.IsConnected(0, index(n-1, m-1, m)) {
					return num
				}
			}

		}
	}
	return 0
}
