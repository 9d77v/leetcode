package main

/*
题目：
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。
进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle-ii/
*/

/*
方法一: 模拟
时间复杂度：О(rowIndex^2)
空间复杂度：О(rowIndex)
运行时间：0 ms	内存消耗：2 MB
*/
func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		result[i] = 1
	}
	for i := 1; i <= rowIndex; i++ {
		tmp := append([]int(nil), result[:i]...)
		for j := 1; j < i; j++ {
			result[j] = tmp[j-1] + tmp[j]
		}
	}
	return result
}

/*
方法二: 模拟
时间复杂度：О(rowIndex^2)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：1.9 MB
*/
func getRowFunc2(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		result[i] = 1
	}
	for i := 1; i <= rowIndex; i++ {
		for j := i - 1; j > 0; j-- {
			result[j] += result[j-1]
		}
	}
	return result
}

/*
方法三: 线性递推
时间复杂度：О(rowIndex)
空间复杂度：О(k)
运行时间：0 ms	内存消耗：1.9 MB
*/
func getRowFunc3(rowIndex int) []int {
	k := rowIndex + 1
	row := make([]int, k)
	row[0] = 1
	for i := 1; i < k; i++ {
		row[i] = row[i-1] * (k - i) / i
	}
	return row
}
