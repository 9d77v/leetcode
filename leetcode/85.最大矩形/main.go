package main

/*
题目：
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积

提示：
rows == matrix.length
cols == matrix[0].length
0 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-rectangle
*/

/*
方法一：双指针 / 从后往前
时间复杂度：О(n+m)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.3 MB
*/
func maximalRectangleFunc1(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	return 0
}
