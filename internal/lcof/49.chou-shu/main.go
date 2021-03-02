package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
说明:

1 是丑数。
n 不超过1690。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/chou-shu-lcof/
*/

/*
方法一：模拟
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：超出时间限制
*/
func nthUglyNumber(n int) int {
	i, x := 0, 0
	for i < n {
		x++
		if isUglyNumber(x) {
			i++
		}
	}
	return x
}

func isUglyNumber(n int) bool {
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}

/*
方法二：模拟
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：4.2 MB
*/
func nthUglyNumberFunc2(n int) int {
	a, b, c, dp := 0, 0, 0, make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = MinArr(n2, n3, n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}
