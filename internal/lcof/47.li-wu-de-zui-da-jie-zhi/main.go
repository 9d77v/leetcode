package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物

提示：

0 < grid.length <= 200
0 < grid[0].length <= 200

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof
*/

/*
方法一：动态规划
f(i,j)=grid[i][j] i=0,j=0
f(i,j)=grid[i][j]+f(i-1,j) i!=0,j=0
f(i,j)=grid[i][j]+f(i,j-1) i=0,j!=0
f(i,j)=grid[i][j]+max(f(i-1,j),f(i,j-1) ) i!=0,j!=0
时间复杂度：О(mn)
空间复杂度：О(mn)
运行时间：8 ms	内存消耗：4.3 MB
*/
func maxValue(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i != 0 && j != 0 {
				dp[i][j] = grid[i][j] + Max(dp[i-1][j], dp[i][j-1])
			} else if i != 0 && j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else if i == 0 && j != 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else {
				dp[i][j] = grid[i][j]
			}
		}
	}
	return dp[n-1][m-1]
}

/*
方法二：动态规划，增加一行一列，避免边界问题
f(i,j)=grid[i][j]+max(f(i-1,j),f(i,j-1) )
时间复杂度：О(mn)
空间复杂度：О(mn)
运行时间：8 ms	内存消耗：4.3 MB
*/
func maxValueFunc2(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = grid[i-1][j-1] + Max(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[n][m]
}
