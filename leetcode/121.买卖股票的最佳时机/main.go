package main

import (
	. "github.com/9d77v/leetcode/lib"
)

/*
题目：
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
*/

/*
股票问题通解
用 n 表示股票价格数组的长度；
用 i 表示第 i 天（i 的取值范围是 0 到 n - 1）；
用 k 表示允许的最大交易次数；
用 T[i][k] 表示在第 i 天结束时，最多进行 k 次交易的情况下可以获得的最大收益。

基准情况：
T[-1][k][0] = 0, T[-1][k][1] = -Infinity
T[i][0][0] = 0, T[i][0][1] = -Infinity
状态转移方程：
T[i][k][0] = max(T[i - 1][k][0], T[i - 1][k][1] + prices[i])
T[i][k][1] = max(T[i - 1][k][1], T[i - 1][k - 1][0] - prices[i])

作者：Storm
链接：https://leetcode-cn.com/circle/article/qiAgHn/
*/

/*
方法一：动态规划
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：4 ms	内存消耗：3.5 MB
*/
func maxProfitFunc1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0] = [2]int{0, -prices[0]}
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

/*
方法二：动态规划
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：3 MB
*/
func maxProfitFunc2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	profit0, profit1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		profit0 = Max(profit0, profit1+prices[i])
		profit1 = Max(profit1, -prices[i])
	}
	return profit0
}
