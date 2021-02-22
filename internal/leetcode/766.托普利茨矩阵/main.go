package main

/*
题目：
给你一个 m x n 的矩阵 matrix 。如果这个矩阵是托普利茨矩阵，返回 true ；否则，返回 false 。

如果矩阵上每一条由左上到右下的对角线上的元素都相同，那么这个矩阵是 托普利茨矩阵

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/toeplitz-matrix
*/

/*
方法一：遍历矩阵
时间复杂度：О(nm)
空间复杂度：О(1)
运行时间：12 ms	内存消耗：4.4 MB
*/
func isToeplitzMatrix(matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])
	for i := n - 2; i > 0; i-- {
		num, k := matrix[i][0], 1
		for j := i + 1; j < n && k < m; j++ {
			if matrix[j][k] != num {
				return false
			}
			k++
		}
	}
	for i := 0; i < m; i++ {
		num, k := matrix[0][i], 1
		for j := i + 1; j < m && k < n; j++ {
			if matrix[k][j] != num {
				return false
			}
			k++
		}
	}
	return true
}

/*
方法二：遍历矩阵
时间复杂度：О(nm)
空间复杂度：О(1)
运行时间：12 ms	内存消耗：4.4 MB
*/
func isToeplitzMatrixFunc2(matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
