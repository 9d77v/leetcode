package main

import "math"

/*
题目：斐波那契数
斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给你 n ，请计算 F(n) 。

提示：
0 <= n <= 30

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fibonacci-number/
*/

/*
方法一：递归
运行时间：12 ms	内存消耗：1.9 MB
*/
func fibFunc1(n int) int {
	if n < 2 {
		return n
	}
	return fibFunc1(n-1) + fibFunc1(n-2)
}

/*
方法二：动态规划
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：1.9 MB
*/
func fibFunc2(n int) int {
	if n < 2 {
		return n
	}
	prev, cur := 0, 1
	for i := 2; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}

/*
方法三：矩阵快速幂
时间复杂度：О(㏒n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：1.9 MB
*/
func fibFunc3(n int) int {
	if n < 2 {
		return n
	}
	res := pow(matrix{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}

type matrix [2][2]int

func multiply(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return
}

func pow(a matrix, n int) matrix {
	res := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = multiply(res, a)
		}
		a = multiply(a, a)
	}
	return res
}

/*
方法四：通项公式
运行时间：0 ms	内存消耗：1.8 MB
*/
func fibFunc4(n int) int {
	sqrt5 := math.Sqrt(5)
	fibN := math.Pow((1+sqrt5)/2, float64(n)) - math.Pow((1-sqrt5)/2, float64(n))
	return int(math.Round(fibN / sqrt5))
}
