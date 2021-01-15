package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。

提示：
1 <= n <= 200
n == isConnected.length
n == isConnected[i].length
isConnected[i][j] 为 1 或 0
isConnected[i][i] == 1
isConnected[i][j] == isConnected[j][i]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-provinces
*/

/*
方法一：深度优先搜索
时间复杂度：О(n^2)
空间复杂度：О(n)
运行时间：36 ms	内存消耗：6.6 MB
*/
func findCircleNumFunc1(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	res := 0
	for i, v := range visited {
		if !v {
			res++
			dfs(isConnected, visited, i)
		}
	}
	return res
}

func dfs(isConnected [][]int, visited []bool, from int) {
	visited[from] = true
	for to, conn := range isConnected[from] {
		if conn == 1 && !visited[to] {
			dfs(isConnected, visited, to)
		}
	}
}

/*
方法二：广度优先搜索
时间复杂度：О(n^2)
空间复杂度：О(n)
运行时间：24 ms	内存消耗：6.7 MB
*/
func findCircleNumFunc2(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	res := 0
	for i, v := range visited {
		if !v {
			res++
			bfs(isConnected, i, visited)
		}
	}
	return res
}

func bfs(isConnected [][]int, from int, visited []bool) {
	queue := []int{from}
	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]
		visited[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !visited[to] {
				queue = append(queue, to)
			}
		}
	}
}

/*
方法三：并查集
时间复杂度：О(n^2logn)
空间复杂度：О(n)
运行时间：36 ms	内存消耗：6.6 MB
*/
func findCircleNumFunc3(isConnected [][]int) int {
	n := len(isConnected)
	var uf UnionFind = NewArrayUnionFind(n)
	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if i == uf.Find(i) {
			res++
		}
	}
	return res
}
