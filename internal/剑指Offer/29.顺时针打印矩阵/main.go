package main

/*
题目：
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
*/

/*
方法一： 模拟
时间复杂度：О(mn)
空间复杂度：О(mn)
运行时间：24 ms	内存消耗：6.2 MB
*/

func spiralOrderFunc1(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}
	columns := len(matrix[0])
	if columns == 0 {
		return []int{}
	}
	result := make([]int, rows*columns)
	start, j := 0, 0
	for rows > start*2 && columns > start*2 {
		j = spiralOrderCore(matrix, columns, rows, result, start, j)
		start++
	}
	return result
}

func spiralOrderCore(matrix [][]int, columns, rows int, result []int, start, j int) int {
	endX, endY := columns-1-start, rows-1-start
	for i := start; i <= endX; i++ {
		result[j] = matrix[start][i]
		j++
	}

	if start < endY {
		for i := start + 1; i <= endY; i++ {
			result[j] = matrix[i][endX]
			j++
		}
	}
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			result[j] = matrix[endY][i]
			j++
		}
	}
	if start < endX && start < endY-1 {
		for i := endY - 1; i >= start+1; i-- {
			result[j] = matrix[i][start]
			j++
		}
	}
	return j
}

/*
方法一： 按层模拟
时间复杂度：О(mn)
空间复杂度：О(mn)
运行时间：24 ms	内存消耗：6.2 MB
*/
func spiralOrderFunc2(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}
	columns := len(matrix[0])
	if columns == 0 {
		return []int{}
	}
	result := make([]int, rows*columns)
	left, right, top, bottom, j := 0, columns-1, 0, rows-1, 0
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result[j] = matrix[top][i]
			j++
		}
		top++
		for i := top; i <= bottom; i++ {
			result[j] = matrix[i][right]
			j++
		}
		right--
		for i := right; i >= left && top <= bottom; i-- {
			result[j] = matrix[bottom][i]
			j++
		}
		bottom--
		for i := bottom; i >= top && left <= right; i-- {
			result[j] = matrix[i][left]
			j++
		}
		left++
	}
	return result
}
