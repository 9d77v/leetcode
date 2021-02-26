package main

/*
题目：转置矩阵
给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。

矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 1000
1 <= m * n <= 10^5
-10^9 <= matrix[i][j] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/transpose-matrix
*/

/*
方法一：复制
时间复杂度：О(mn)
空间复杂度：О(mn)
运行时间：8 ms	内存消耗：6.1 MB
*/
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}
