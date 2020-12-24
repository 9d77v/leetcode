package main

/*
题目：在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
限制：
0 <= n <= 1000
0 <= m <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
*/

/*
方法一：暴力求解
时间复杂度：О(nm)
空间复杂度：О(1)
运行时间：24 ms	内存消耗：6.6 MB
*/
func findNumberIn2DArrayFunc1(matrix [][]int, target int) bool {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

/*
方法二：线性查找
时间复杂度：О(n+m)
空间复杂度：О(1)
运行时间：28 ms	内存消耗：6.6 MB
*/
func findNumberIn2DArrayFunc2(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	row, column := 0, len(matrix[0])-1
	for row < rows && column >= 0 {
		num := matrix[row][column]
		if num == target {
			return true
		}
		if num > target {
			column--
		} else {
			row++
		}
	}
	return false
}
