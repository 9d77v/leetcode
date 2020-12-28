package main

import (
	. "github.com/9d77v/leetcode/lib"
)

/*
题目：
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

提示：
1 <= prices.length <= 3 * 10 ^ 4
0 <= prices[i] <= 10 ^ 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
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
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

/*
方法二：动态规划
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：3 MB
*/
func maxProfitFunc2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	profit0, profit1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		profit0 = Max(profit0, profit1+prices[i])
		profit1 = Max(profit1, profit0-prices[i])
	}
	return profit0
}
