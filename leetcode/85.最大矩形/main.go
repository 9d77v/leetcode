package main

import (
	. "github.com/9d77v/leetcode/lib"
)

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
方法一：单调栈
时间复杂度：О(m²n)
空间复杂度：О(mn)
运行时间：4 ms	内存消耗：6.1 MB
*/
func maximalRectangleFunc1(matrix [][]byte) (max int) {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	heights := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != '0' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		max = Max(max, largestRectangleArea(heights))
	}
	return
}

//最大矩形面积
func largestRectangleArea(heights []int) (max int) {
	//前后补0
	heights = append([]int{0}, append(heights, 0)...)
	n := len(heights)
	stack := NewStack(n)
	for i := 0; i < n; i++ {
		for stack.IsNotEmpty() && heights[stack.Top().(int)] > heights[i] {
			cur := stack.Pop().(int)
			left := stack.Top().(int) + 1
			right := i - 1
			max = Max(max, (right-left+1)*heights[cur])
		}
		stack.Push(i)
	}
	return
}
