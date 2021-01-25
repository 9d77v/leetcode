package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
在由 1 x 1 方格组成的 N x N 网格 grid 中，每个 1 x 1 方块由 /、\ 或空格构成。这些字符会将方块划分为一些共边的区域。

（请注意，反斜杠字符是转义的，因此 \ 用 "\\" 表示。）。

返回区域的数目。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regions-cut-by-slashes
*/

/*
方法一：并查集
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：3.5 MB
     0
  3     1
	 2
*/
func regionsBySlashes(grid []string) int {
	n := len(grid)
	var uf UnionFind = NewArrayUnionFind(4 * n * n)
	for i := range grid {
		for j := range grid {
			num := (i*n + j) * 4
			if i > 0 {
				topNum := ((i-1)*n + j) * 4
				uf.Union(topNum+2, num)
			}
			if j > 0 {
				leftNum := (i*n + j - 1) * 4
				uf.Union(leftNum+1, num+3)
			}
			switch grid[i][j] {
			case ' ':
				uf.Union(num, num+1)
				uf.Union(num, num+2)
				uf.Union(num, num+3)
			case '/':
				uf.Union(num, num+3)
				uf.Union(num+1, num+2)
			case '\\':
				uf.Union(num, num+1)
				uf.Union(num+2, num+3)
			}
		}
	}
	return uf.Count()
}
