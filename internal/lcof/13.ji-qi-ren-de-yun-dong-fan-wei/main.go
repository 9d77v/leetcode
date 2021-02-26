package main

/*
题目：机器人的运动范围
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

提示：
1 <= n,m <= 100
0 <= k <= 20

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
*/

/*
方法一：回溯法
时间复杂度：О(mn)
空间复杂度：О(mn)
运行时间：0 ms	内存消耗：3 MB
*/
func movingCount(m int, n int, k int) int {
	board := make([][]bool, m)
	for i := range board {
		board[i] = make([]bool, n)
	}
	return dfs(board, 0, 0, k)
}

func dfs(board [][]bool, i, j, k int) int {
	count := 0
	if check(board, i, j, k) {
		board[i][j] = true
		count = 1 + dfs(board, i+1, j, k) + dfs(board, i, j+1, k) + dfs(board, i-1, j, k) + dfs(board, i, j-1, k)
	}
	return count
}

func check(board [][]bool, i, j, k int) bool {
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) && getDigitalSum(i)+getDigitalSum(j) <= k && !board[i][j] {
		return true
	}
	return false
}

func getDigitalSum(i int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i /= 10
	}
	return sum
}
