package main

/*
题目：重塑矩阵
在MATLAB中，有一个非常有用的函数 reshape，它可以将一个矩阵重塑为另一个大小不同的新矩阵，但保留其原始数据。

给出一个由二维数组表示的矩阵，以及两个正整数r和c，分别表示想要的重构的矩阵的行数和列数。

重构后的矩阵需要将原始矩阵的所有元素以相同的行遍历顺序填充。

如果具有给定参数的reshape操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。

注意：

给定矩阵的宽和高范围在 [1, 100]。
给定的 r 和 c 都是正数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reshape-the-matrix
*/

/*
方法一：遍历
时间复杂度：О(rc)
空间复杂度：О(1)
运行时间：16 ms	内存消耗：6.9 MB
*/
func matrixReshape(nums [][]int, r int, c int) [][]int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	m := len(nums[0])
	if m == 0 {
		return nums
	}
	if n*m != r*c {
		return nums
	}
	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}
	for i := 0; i < r*c; i++ {
		result[i/c][i%c] = nums[i/m][i%m]
	}
	return result
}
