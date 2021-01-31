package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
有一个 m x n 的二元网格，其中 1 表示砖块，0 表示空白。砖块 稳定（不会掉落）的前提是：

一块砖直接连接到网格的顶部，或者
至少有一块相邻（4 个方向之一）砖块 稳定 不会掉落时
给你一个数组 hits ，这是需要依次消除砖块的位置。每当消除 hits[i] = (rowi, coli) 位置上的砖块时，对应位置的砖块（若存在）会消失，然后其他的砖块可能因为这一消除操作而掉落。一旦砖块掉落，它会立即从网格中消失（即，它不会落在其他稳定的砖块上）。

返回一个数组 result ，其中 result[i] 表示第 i 次消除操作对应掉落的砖块数目。

注意，消除可能指向是没有砖块的空白位置，如果发生这种情况，则没有砖块掉落。

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
grid[i][j] 为 0 或 1
1 <= hits.length <= 4 * 104
hits[i].length == 2
0 <= xi <= m - 1
0 <= yi <= n - 1
所有 (xi, yi) 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bricks-falling-when-hit
*/
func hitBricks(grid [][]int, hits [][]int) []int {
	m, n := len(grid), len(grid[0])
	status := make([][]int, m)
	for i, row := range grid {
		status[i] = append([]int(nil), row...)
	}
	for _, p := range hits {
		status[p[0]][p[1]] = 0
	}

	var uf UnionFind = NewArrayUnionFind(m*n + 1)
	root := m * n
	for i, row := range status {
		for j, v := range row {
			if v == 0 {
				continue
			}
			if i == 0 {
				uf.Union(i*n+j, root)
			}
			if i > 0 && status[i-1][j] == 1 {
				uf.Union(i*n+j, (i-1)*n+j)
			}
			if j > 0 && status[i][j-1] == 1 {
				uf.Union(i*n+j, i*n+j-1)
			}
		}
	}

	type pair struct{ x, y int }
	directions := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

	res := make([]int, len(hits))
	for i := len(hits) - 1; i >= 0; i-- {
		p := hits[i]
		r, c := p[0], p[1]
		if grid[r][c] == 0 {
			continue
		}

		preSize := uf.Rank(uf.Find(root))
		if r == 0 {
			uf.Union(c, root)
		}
		for _, d := range directions {
			if newR, newC := r+d.x, c+d.y; 0 <= newR && newR < m && 0 <= newC && newC < n && status[newR][newC] == 1 {
				uf.Union(r*n+c, newR*n+newC)
			}
		}
		curSize := uf.Rank(uf.Find(root))
		if cnt := curSize - preSize - 1; cnt > 0 {
			res[i] = cnt
		}
		status[r][c] = 1
	}
	return res
}
