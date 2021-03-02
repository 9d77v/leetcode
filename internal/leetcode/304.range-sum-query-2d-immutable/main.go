package main

/*
题目：
给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2)。
matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。

说明:

你可以假设矩阵不可变。
会多次调用 sumRegion 方法。
你可以假设 row1 ≤ row2 且 col1 ≤ col2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-query-2d-immutable
*/

/*
方法一: 一维前缀和
初始化计算前缀和：preSums[i][j+1] = preSums[i][j] + num
对于（row1,col1) (row2,col2)
遍历row1 <=i<=row2
sum+=preSum[i][col2+1]-preSum[i][col1]
时间复杂度：О(mn)，初始化O(mn),检索O（m)
空间复杂度：О(mn)
运行时间：40 ms	内存消耗：10.6 MB
*/

type NumMatrix struct {
	preSums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	preSums := make([][]int, len(matrix))
	for i, row := range matrix {
		preSums[i] = make([]int, len(row)+1)
		for j, num := range row {
			preSums[i][j+1] = preSums[i][j] + num
		}
	}
	return NumMatrix{preSums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (sum int) {
	for i := row1; i <= row2; i++ {
		sum += this.preSums[i][col2+1] - this.preSums[i][col1]
	}
	return
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

/*
方法二: 二维前缀和
preSums[i][j] = preSums[i][j-1] + preSums[i-1][j]-preSum[i-1][j-1] +num
对于（row1,col1) (row2,col2)
sum=preSums[row2][col2]-preSums[row1-1][col2]-preSums[row2][col1-1]+preSums[row1-1][col1-1]
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：40 ms	内存消耗：10.6 MB
*/

type NumMatrix2 struct {
	preSums [][]int
}

func Constructor2(matrix [][]int) NumMatrix2 {
	if len(matrix) == 0 {
		return NumMatrix2{}
	}
	preSums := make([][]int, len(matrix)+1)
	preSums[0] = make([]int, len(matrix[0])+1)
	for i, row := range matrix {
		preSums[i+1] = make([]int, len(row)+1)
		for j, num := range row {
			preSums[i+1][j+1] = preSums[i+1][j] + preSums[i][j+1] - preSums[i][j] + num
		}
	}
	return NumMatrix2{preSums}
}

func (this *NumMatrix2) SumRegion(row1 int, col1 int, row2 int, col2 int) (sum int) {
	return this.preSums[row2+1][col2+1] - this.preSums[row1][col2+1] - this.preSums[row2+1][col1] + this.preSums[row1][col1]
}
